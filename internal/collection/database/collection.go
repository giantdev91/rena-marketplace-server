package database

import (
	"context"
	"fmt"
	"rena-server-v2/internal/collection/model"
	"rena-server-v2/internal/database"
	"rena-server-v2/pkg/logging"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IterateCollectionCriteria struct {
	Slug   string
	Offset uint
	Limit  uint
}

//go:generate mockery --name CollectionDB --filename collection_mock.go
type CollectionDB interface {
	RunInTx(ctx context.Context, f func(ctx context.Context) error) error

	// SaveCollection saves a given collection.
	SaveCollection(ctx context.Context, collection *model.Collection) error

	// FindCollectionBySlug returns a collection with given slug
	// database.ErrNotFound error is returned if not exist
	FindCollectionBySlug(ctx context.Context, slug string) (*model.Collection, error)

	// FindCollectionByAddress returns a collection with give contract address
	// and returns error if not exist
	FindCollectionByAddress(ctx context.Context, address string) (*model.Collection, error)

	// FindCollections returns collection list with given criteria and total count
	FindCollections(ctx context.Context, criteria IterateCollectionCriteria) ([]*model.Collection, int64, error)

	// updateCollection changes a given collection.
	UpdateCollection(ctx context.Context, collection *model.Collection) error

	// DeleteCollectionBySlug deletes a collection with given slug
	// and returns nil if success to delete, otherwise returns an error
	DeleteCollectionBySlug(ctx context.Context, slug string) error
}

type collectionDB struct {
	db *gorm.DB
}

func (a *collectionDB) RunInTx(ctx context.Context, f func(ctx context.Context) error) error {
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

func (a *collectionDB) SaveCollection(ctx context.Context, collection *model.Collection) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("collection.db.SaveCollection", "collection", collection)

	if err := db.WithContext(ctx).Create(collection).Error; err != nil {
		logger.Errorw("collection.db.SaveCollection failed to save collection", "err", err)
		if database.IsKeyConflictErr(err) {
			return database.ErrKeyConflict
		}
		return err
	}
	return nil
}

func (a *collectionDB) FindCollectionBySlug(ctx context.Context, slug string) (*model.Collection, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("collection.db.FindCollectionBySlug", "slug", slug)

	var ret model.Collection
	// SELECT collections.*
	// FROM `collections`
	// WHERE slug = "title1" ORDER BY `collections`.`id` LIMIT 1
	err := db.WithContext(ctx).First(&ret, "slug = ?", slug).Error

	if err != nil {
		logger.Errorw("failed to find collection", "err", err)
		if database.IsRecordNotFoundErr(err) {
			return nil, database.ErrNotFound
		}
		return nil, err
	}
	return &ret, nil
}

func (a *collectionDB) FindCollectionByAddress(ctx context.Context, address string) (*model.Collection, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("collection.db.FindCollectionByAddress", "address", address)

	var ret model.Collection
	// SELECT collections.*
	// FROM `collections`
	// WHERE contract_address = "0x..." ORDER BY `collections`.`id` LIMIT 1
	err := db.WithContext(ctx).First(&ret, "contract_address = ?", address).Error

	if err != nil {
		logger.Errorw("failed to find collection", "err", err)
		if database.IsRecordNotFoundErr(err) {
			return nil, database.ErrNotFound
		}
		return nil, err
	}
	return &ret, nil
}

func (a *collectionDB) FindCollections(ctx context.Context, criteria IterateCollectionCriteria) ([]*model.Collection, int64, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("collection.db.FindCollections", "criteria", criteria)

	chain := db.WithContext(ctx).Table("collections a")

	// get total count
	var totalCount int64
	err := chain.Select("a.id").Count(&totalCount).Error
	if err != nil {
		logger.Error("failed to get total count", "err", err)
	}

	// get collection ids
	rows, err := chain.Select("DISTINCT(a.id) id").
		Offset(int(criteria.Offset)).
		Limit(int(criteria.Limit)).
		Order("a.id DESC").
		Rows()
	if err != nil {
		logger.Error("failed to read collection ids", "err", err)
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

	// get collections by ids
	var ret []*model.Collection
	if len(ids) == 0 {
		return []*model.Collection{}, totalCount, nil
	}
	err = db.WithContext(ctx).
		Where("collections.id IN (?)", ids).
		Order("collections.id DESC").
		Find(&ret).Error
	if err != nil {
		logger.Error("failed to find collection by ids", "err", err)
		return nil, 0, err
	}

	return ret, totalCount, nil
}

func (a *collectionDB) UpdateCollection(ctx context.Context, collection *model.Collection) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("collection.db.Update", "collection", collection)

	chain := db.WithContext(ctx).
		Model(&model.Collection{}).
		Where("contract_address = ?", collection.ContractAddress).
		UpdateColumns(collection)
	if chain.Error != nil {
		logger.Error("account.db.Update failed to update", "err", chain.Error)
		return chain.Error
	}
	if chain.RowsAffected == 0 {
		return database.ErrNotFound
	}
	return nil
}

func (a *collectionDB) DeleteCollectionBySlug(ctx context.Context, slug string) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("collection.db.DeleteCollectionBySlug", "slug", slug)

	// delete collection
	chain := db.WithContext(ctx).Where("slug = ?", slug).Delete(&model.Collection{})

	if chain.Error != nil {
		logger.Errorw("failed to delete an collection", "err", chain.Error)
		return chain.Error
	}
	if chain.RowsAffected == 0 {
		logger.Error("failed to delete an collection because not found")
		return database.ErrNotFound
	}
	return nil
}

// NewCollectionDB creates a new collection db with given db
func NewCollectionDB(db *gorm.DB) CollectionDB {
	return &collectionDB{
		db: db,
	}
}
