package database

import (
	"context"
	"fmt"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/registry/model"
	"rena-server-v2/pkg/logging"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IterateRegistryCriteria struct {
	Type        string
	SortBy      string
	FilterBy    string
	Category    string
	Address     string
	Collections []uint
	PriceMin    float32
	PriceMax    float32
	Payment     string
	Offset      uint
	Limit       uint
}

//go:generate mockery --name RegistryDB --filename registry_mock.go
type RegistryDB interface {
	RunInTx(ctx context.Context, f func(ctx context.Context) error) error

	// SaveRegistry saves a given registry.
	SaveRegistry(ctx context.Context, registry *model.Registry) error

	// Update updates a given registry
	UpdateRegistry(ctx context.Context, contractAddress string, tokenID uint, payToken string, newPrice uint64, registry *model.Registry) error

	// FindRegistries returns registry list with given criteria and total count
	FindRegistries(ctx context.Context, criteria IterateRegistryCriteria) ([]*model.Registry, int64, error)

	// DeleteRegistryBySlug deletes a registry with given slug
	// and returns nil if success to delete, otherwise returns an error
	DeleteRegistryBySlug(ctx context.Context, authorId uint, slug string) error
}

type registryDB struct {
	db *gorm.DB
}

func (a *registryDB) RunInTx(ctx context.Context, f func(ctx context.Context) error) error {
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

func (a *registryDB) SaveRegistry(ctx context.Context, registry *model.Registry) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("registry.db.SaveRegistry", "registry", registry)

	if err := db.WithContext(ctx).Create(registry).Error; err != nil {
		logger.Errorw("registry.db.SaveRegistry failed to save registry", "err", err)
		if database.IsKeyConflictErr(err) {
			return database.ErrKeyConflict
		}
		return err
	}
	return nil
}

func (a *registryDB) UpdateRegistry(ctx context.Context, contractAddress string, tokenID uint, payToken string, newPrice uint64, registry *model.Registry) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("registry.db.Update", "registry", registry)

	fields := make(map[string]interface{})
	fields["contract_address"] = contractAddress
	fields["token_id"] = tokenID

	chain := db.WithContext(ctx).
		Model(&model.Registry{}).
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

func (a *registryDB) FindRegistries(ctx context.Context, criteria IterateRegistryCriteria) ([]*model.Registry, int64, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("registry.db.FindRegistries", "criteria", criteria)

	chain := db.WithContext(ctx).Table("registrys a").Where("deleted_at_unix = 0")
	if len(criteria.Collections) != 0 {
		chain = chain.Where("a.collection_id IN ?", criteria.Collections)
	}
	if criteria.Payment == "eth" {
		chain = chain.Where("a.price > ? AND a.price < ?", criteria.PriceMin, criteria.PriceMax)
	}
	if criteria.Payment == "usd" {
		chain = chain.Where("a.price_in_usd > ? AND a.price_in_usd < ?", criteria.PriceMin, criteria.PriceMax)
	}

	// get total count
	var totalCount int64
	err := chain.Select("a.id").Count(&totalCount).Error
	if err != nil {
		logger.Error("failed to get total count", "err", err)
	}

	// get registry ids
	rows, err := chain.Select("DISTINCT(a.id) id").
		Offset(int(criteria.Offset)).
		Limit(int(criteria.Limit)).
		Order("a.id DESC").
		Rows()
	if err != nil {
		logger.Error("failed to read registry ids", "err", err)
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

	// get registrys by ids
	var ret []*model.Registry
	if len(ids) == 0 {
		return []*model.Registry{}, totalCount, nil
	}
	err = db.WithContext(ctx).Joins("Collection").
		Where("registrys.id IN (?)", ids).
		Order("registrys.id DESC").
		Find(&ret).Error
	if err != nil {
		logger.Error("failed to find registry by ids", "err", err)
		return nil, 0, err
	}
	fmt.Println(ret)

	return ret, totalCount, nil
}

func (a *registryDB) DeleteRegistryBySlug(ctx context.Context, authorId uint, slug string) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("registry.db.DeleteRegistryBySlug", "slug", slug)

	// delete registry
	chain := db.WithContext(ctx).Model(&model.Registry{}).
		Where("slug = ? AND deleted_at_unix = 0", slug).
		Where("author_id = ?", authorId).
		Update("deleted_at_unix", time.Now().Unix())
	if chain.Error != nil {
		logger.Errorw("failed to delete an registry", "err", chain.Error)
		return chain.Error
	}
	if chain.RowsAffected == 0 {
		logger.Error("failed to delete an registry because not found")
		return database.ErrNotFound
	}
	return nil
}

// NewRegistryDB creates a new registry db with given db
func NewRegistryDB(db *gorm.DB) RegistryDB {
	return &registryDB{
		db: db,
	}
}
