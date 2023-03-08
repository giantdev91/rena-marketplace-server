package database

import (
	"context"
	"fmt"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/marketplace/model"
	"rena-server-v2/pkg/logging"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IterateMarketplaceCriteria struct {
	Type   string
	SortBy string
	Offset uint
	Limit  uint
}

//go:generate mockery --name MarketplaceDB --filename marketplace.go
type MarketplaceDB interface {
	RunInTx(ctx context.Context, f func(ctx context.Context) error) error

	// SaveMarketplace saves a given marketplace.
	SaveMarketplace(ctx context.Context, marketplace *model.Marketplace) error

	// Update updates a given marketplace
	UpdateMarketplace(ctx context.Context, contractAddress string, tokenID uint, payToken string, newPrice uint64, marketplace *model.Marketplace) error

	// FindMarketplaceBySlug returns a marketplace with given slug
	// database.ErrNotFound error is returned if not exist
	FindMarketplaceBySlug(ctx context.Context, slug string) (*model.Marketplace, error)

	// FindMarketplaces returns marketplace list with given criteria and total count
	FindMarketplaces(ctx context.Context, criteria IterateMarketplaceCriteria) ([]*model.Marketplace, int64, error)
}

type marketplaceDB struct {
	db *gorm.DB
}

func (a *marketplaceDB) RunInTx(ctx context.Context, f func(ctx context.Context) error) error {
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

func (a *marketplaceDB) SaveMarketplace(ctx context.Context, marketplace *model.Marketplace) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("marketplace.db.SaveMarketplace", "marketplace", marketplace)

	if err := db.WithContext(ctx).Create(marketplace).Error; err != nil {
		logger.Errorw("marketplace.db.SaveMarketplace failed to save marketplace", "err", err)
		if database.IsKeyConflictErr(err) {
			return database.ErrKeyConflict
		}
		return err
	}
	return nil
}

func (a *marketplaceDB) UpdateMarketplace(ctx context.Context, contractAddress string, tokenID uint, payToken string, newPrice uint64, marketplace *model.Marketplace) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("marketplace.db.Update", "marketplace", marketplace)

	fields := make(map[string]interface{})
	fields["contract_address"] = contractAddress
	fields["token_id"] = tokenID

	chain := db.WithContext(ctx).
		Model(&model.Marketplace{}).
		Where("contract_address = ? AND token_id = ?", contractAddress, tokenID).
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

func (a *marketplaceDB) FindMarketplaceBySlug(ctx context.Context, slug string) (*model.Marketplace, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("marketplace.db.FindMarketplaceBySlug", "slug", slug)

	var ret model.Marketplace
	// 1) load marketplace with author
	// SELECT marketplaces.*, accounts.*
	// FROM `marketplaces` LEFT JOIN `accounts` `Author` ON `marketplaces`.`author_id` = `Author`.`id`
	// WHERE slug = "title1" AND deleted_at_unix = 0 ORDER BY `marketplaces`.`id` LIMIT 1
	err := db.WithContext(ctx).Joins("Author").
		First(&ret, "slug = ? AND deleted_at_unix = 0", slug).Error

	if err != nil {
		logger.Errorw("failed to find marketplace", "err", err)
		if database.IsRecordNotFoundErr(err) {
			return nil, database.ErrNotFound
		}
		return nil, err
	}
	return &ret, nil
}

func (a *marketplaceDB) FindMarketplaces(ctx context.Context, criteria IterateMarketplaceCriteria) ([]*model.Marketplace, int64, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("marketplace.db.FindMarketplaces", "criteria", criteria)

	chain := db.WithContext(ctx).Table("marketplaces a").Where("deleted_at_unix = 0")

	// get total count
	var totalCount int64
	err := chain.Select("a.id").Count(&totalCount).Error
	if err != nil {
		logger.Error("failed to get total count", "err", err)
	}

	// get marketplace ids
	rows, err := chain.Select("DISTINCT(a.id) id").
		Offset(int(criteria.Offset)).
		Limit(int(criteria.Limit)).
		Order("a.id DESC").
		Rows()
	if err != nil {
		logger.Error("failed to read marketplace ids", "err", err)
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

	// get marketplaces by ids
	var ret []*model.Marketplace
	if len(ids) == 0 {
		return []*model.Marketplace{}, totalCount, nil
	}
	err = db.WithContext(ctx).Joins("Collection").
		Where("marketplaces.id IN (?)", ids).
		Order("marketplaces.id DESC").
		Find(&ret).Error
	if err != nil {
		logger.Error("failed to find marketplace by ids", "err", err)
		return nil, 0, err
	}
	fmt.Println(ret)

	return ret, totalCount, nil
}

// NewMarketplaceDB creates a new marketplace db with given db
func NewMarketplaceDB(db *gorm.DB) MarketplaceDB {
	return &marketplaceDB{
		db: db,
	}
}
