package database

import (
	"context"
	"fmt"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/token/model"
	"rena-server-v2/pkg/logging"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IterateTokenCriteria struct {
	Offset uint
	Limit  uint
}

//go:generate mockery --name TokenDB --filename token.go
type TokenDB interface {
	RunInTx(ctx context.Context, f func(ctx context.Context) error) error

	// SaveToken saves a given token.
	SaveToken(ctx context.Context, token *model.Token) error

	// FindTokens returns token list with given criteria and total count
	FindTokens(ctx context.Context, criteria IterateTokenCriteria) ([]*model.Token, int64, error)

	// DeleteToken deletes a token with given slug
	// and returns nil if success to delete, otherwise returns an error
	DeleteToken(ctx context.Context, address string) error
}

type tokenDB struct {
	db *gorm.DB
}

func (a *tokenDB) RunInTx(ctx context.Context, f func(ctx context.Context) error) error {
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

func (a *tokenDB) SaveToken(ctx context.Context, token *model.Token) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("token.db.SaveToken", "token", token)

	if err := db.WithContext(ctx).Create(token).Error; err != nil {
		logger.Errorw("token.db.SaveToken failed to save token", "err", err)
		if database.IsKeyConflictErr(err) {
			return database.ErrKeyConflict
		}
		return err
	}
	return nil
}

func (a *tokenDB) FindTokens(ctx context.Context, criteria IterateTokenCriteria) ([]*model.Token, int64, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("token.db.FindTokens", "criteria", criteria)

	chain := db.WithContext(ctx).Table("tokens a").Where("deleted_at_unix = 0")

	// get total count
	var totalCount int64
	err := chain.Select("a.id").Count(&totalCount).Error
	if err != nil {
		logger.Error("failed to get total count", "err", err)
	}

	// get token ids
	rows, err := chain.Select("DISTINCT(a.id) id").
		Offset(int(criteria.Offset)).
		Limit(int(criteria.Limit)).
		Order("a.id DESC").
		Rows()
	if err != nil {
		logger.Error("failed to read token ids", "err", err)
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

	// get tokens by ids
	var ret []*model.Token
	if len(ids) == 0 {
		return []*model.Token{}, totalCount, nil
	}
	err = db.WithContext(ctx).
		Where("tokens.id IN (?)", ids).
		Order("tokens.id DESC").
		Find(&ret).Error
	if err != nil {
		logger.Error("failed to find token by ids", "err", err)
		return nil, 0, err
	}
	fmt.Println(ret)

	return ret, totalCount, nil
}

func (a *tokenDB) DeleteToken(ctx context.Context, address string) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("token.db.DeleteToken", "address", address)

	// delete token
	chain := db.WithContext(ctx).Model(&model.Token{}).
		Where("token = ? AND deleted_at_unix = 0", address).
		Update("deleted_at_unix", time.Now().Unix())
	if chain.Error != nil {
		logger.Errorw("failed to delete an token", "err", chain.Error)
		return chain.Error
	}
	if chain.RowsAffected == 0 {
		logger.Error("failed to delete an token because not found")
		return database.ErrNotFound
	}
	return nil
}

// NewTokenDB creates a new token db with given db
func NewTokenDB(db *gorm.DB) TokenDB {
	return &tokenDB{
		db: db,
	}
}
