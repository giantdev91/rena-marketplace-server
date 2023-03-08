package database

import (
	"context"
	"fmt"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/item/model"
	"rena-server-v2/pkg/logging"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IterateItemCriteria struct {
	Type        string
	SortBy      string
	FilterBy    string
	Category    string
	Address     string
	Creator     string
	Owner       string
	Collections []uint
	PriceMin    uint64
	PriceMax    uint64
	Offset      uint
	Limit       uint
}

//go:generate mockery --name ItemDB --filename item_mock.go
type ItemDB interface {
	RunInTx(ctx context.Context, f func(ctx context.Context) error) error

	// SaveItem saves a given item.
	SaveItem(ctx context.Context, item *model.Item) error

	// Update updates a given item
	UpdateItem(ctx context.Context, item *model.Item) error

	// FindItemBySlug returns a item with given slug
	// database.ErrNotFound error is returned if not exist
	FindItemBySlug(ctx context.Context, slug string) (*model.Item, error)

	// FindItemByAddress returns a item with given address
	// database.ErrNotFound error is returned if not exist
	FindItemByAddress(ctx context.Context, address string, id uint) (*model.Item, error)

	// FindItems returns item list with given criteria and total count
	FindItems(ctx context.Context, criteria IterateItemCriteria) ([]*model.Item, int64, error)

	// DeleteItemBySlug deletes a item with given slug
	// and returns nil if success to delete, otherwise returns an error
	DeleteItemBySlug(ctx context.Context, authorId uint, slug string) error
}

type itemDB struct {
	db *gorm.DB
}

func (a *itemDB) RunInTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx := a.db.Begin()
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "start tx")
	}

	ctx = database.WithDB(ctx, tx)
	if err := f(ctx); err != nil {
		if err1 := tx.Rollback().Error; err1 != nil {
			return errors.Wrap(err, fmt.Sprintf("rollback tx: %v", err1.Error()))
		}
		return errors.Wrap(err, "invoke function")
	}
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit tx: %v", err)
	}
	return nil
}

func (a *itemDB) SaveItem(ctx context.Context, item *model.Item) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("item.db.SaveItem", "item", item)

	if err := db.WithContext(ctx).Create(item).Error; err != nil {
		logger.Errorw("item.db.SaveItem failed to save item", "err", err)
		if database.IsKeyConflictErr(err) {
			return database.ErrKeyConflict
		}
		return err
	}
	return nil
}

func (a *itemDB) UpdateItem(ctx context.Context, item *model.Item) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("item.db.Update", "item", item)

	fields := make(map[string]interface{})
	fields["payment_token"] = item.PaymentToken
	fields["price"] = item.Price
	fields["supply"] = item.Supply

	chain := db.WithContext(ctx).
		Model(&model.Item{}).
		Where("contract_address = ? AND token_id = ?", item.ContractAddress, item.TokenID).
		UpdateColumns(fields)
	if chain.Error != nil {
		logger.Error("account.db.Update failed to update", "err", chain.Error)
		return chain.Error
	}
	if chain.RowsAffected == 0 {
		return database.ErrNotFound
	}
	return nil
}

func (a *itemDB) FindItemBySlug(ctx context.Context, slug string) (*model.Item, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("item.db.FindItemBySlug", "slug", slug)

	var ret model.Item
	err := db.WithContext(ctx).Joins("Author").
		First(&ret, "slug = ? AND deleted_at_unix = 0", slug).Error

	if err != nil {
		logger.Errorw("failed to find item", "err", err)
		if database.IsRecordNotFoundErr(err) {
			return nil, database.ErrNotFound
		}
		return nil, err
	}
	return &ret, nil
}

func (a *itemDB) FindItemByAddress(ctx context.Context, address string, id uint) (*model.Item, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("item.db.FindItemByAddress", "address", address, "id", id)

	var ret model.Item
	err := db.First(&ret, "contract_address = ? AND token_id = ?", address, id).Error

	if err != nil {
		if database.IsRecordNotFoundErr(err) {
			return nil, database.ErrNotFound
		}
		return nil, err
	}
	return &ret, nil
}

func (a *itemDB) FindItems(ctx context.Context, criteria IterateItemCriteria) ([]*model.Item, int64, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("item.db.FindItems", "criteria", criteria)

	chain := db.WithContext(ctx).Table("items a").Where("deleted_at_unix = 0")
	if len(criteria.Collections) != 0 {
		chain = chain.Where("a.collection_id IN ?", criteria.Collections)
	}
	if criteria.Creator != "" {
		chain = chain.Where("a.creator = ?", criteria.Creator)
	}
	if criteria.Owner != "" {
		chain = chain.Where("a.owner = ?", criteria.Owner)
	}
	if criteria.PriceMin != 0 {
		chain = chain.Where("a.price > ?", criteria.PriceMin)
	}
	if criteria.PriceMax != 0 {
		chain = chain.Where("a.price < ?", criteria.PriceMax)
	}

	// get total count
	var totalCount int64
	err := chain.Select("a.id").Count(&totalCount).Error
	if err != nil {
		logger.Error("failed to get total count", "err", err)
	}

	// get item ids
	rows, err := chain.Select("DISTINCT(a.id) id").
		Offset(int(criteria.Offset)).
		Limit(int(criteria.Limit)).
		Order("a.id DESC").
		Rows()
	if err != nil {
		logger.Error("failed to read item ids", "err", err)
		return nil, 0, err
	}
	var ids []uint
	for rows.Next() {
		var id uint
		err := rows.Scan(&id)
		if err != nil {
			logger.Error("failed to scan id from id rows", "err", err)
			return nil, 0, err
		}
		ids = append(ids, id)
	}

	// get items by ids
	var ret []*model.Item
	if len(ids) == 0 {
		return []*model.Item{}, totalCount, nil
	}
	err = db.WithContext(ctx).Joins("Collection").
		Where("items.id IN (?)", ids).
		Order("items.id DESC").
		Find(&ret).Error
	if err != nil {
		logger.Error("failed to find item by ids", "err", err)
		return nil, 0, err
	}
	fmt.Println(ret)

	return ret, totalCount, nil
}

func (a *itemDB) DeleteItemBySlug(ctx context.Context, authorId uint, slug string) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("item.db.DeleteItemBySlug", "slug", slug)

	// delete item
	chain := db.WithContext(ctx).Model(&model.Item{}).
		Where("slug = ? AND deleted_at_unix = 0", slug).
		Where("author_id = ?", authorId).
		Update("deleted_at_unix", time.Now().Unix())
	if chain.Error != nil {
		logger.Errorw("failed to delete an item", "err", chain.Error)
		return chain.Error
	}
	if chain.RowsAffected == 0 {
		logger.Error("failed to delete an item because not found")
		return database.ErrNotFound
	}
	return nil
}

// NewItemDB creates a new item db with given db
func NewItemDB(db *gorm.DB) ItemDB {
	return &itemDB{
		db: db,
	}
}
