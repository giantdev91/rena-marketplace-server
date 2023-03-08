package database

import (
	"context"
	"fmt"
	"rena-server-v2/internal/account/model"
	"rena-server-v2/internal/database"
	"rena-server-v2/pkg/logging"

	"gorm.io/gorm"
)

type IterateUserCriteria struct {
	Address string
	Offset  uint
	Limit   uint
}

//go:generate mockery --name AccountDB --filename account_mock.go
type AccountDB interface {
	// Save saves a given account
	Save(ctx context.Context, account *model.Account) error

	// Update updates a given account
	Update(ctx context.Context, address string, account *model.Account) error

	// FindByAddress returns an account with given address if exist
	FindByAddress(ctx context.Context, address string) (*model.Account, error)

	// FindUsers returns user list with given criteria and total count
	FindUsers(ctx context.Context, criteria IterateUserCriteria) ([]*model.Account, int64, error)
}

type accountDB struct {
	db *gorm.DB
}

func (a *accountDB) Save(ctx context.Context, account *model.Account) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("account.db.Save", "account", account)

	if err := db.WithContext(ctx).Create(account).Error; err != nil {
		logger.Error("account.db.Save failed to save", "err", err)
		if database.IsKeyConflictErr(err) {
			return database.ErrKeyConflict
		}
		return err
	}
	return nil
}

func (a *accountDB) Update(ctx context.Context, address string, account *model.Account) error {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("account.db.Update", "account", account)

	fields := make(map[string]interface{})
	if account.Address != "" {
		fields["address"] = account.Address
	}
	if account.Username != "" {
		fields["username"] = account.Username
	}
	if account.Bio != "" {
		fields["bio"] = account.Bio
	}
	if account.Image != "" {
		fields["image"] = account.Image
	}

	chain := db.WithContext(ctx).
		Model(&model.Account{}).
		Where("address = ?", address).
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

func (a *accountDB) FindUsers(ctx context.Context, criteria IterateUserCriteria) ([]*model.Account, int64, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("account.db.FindUsers", "criteria", criteria)

	chain := db.WithContext(ctx).Table("accounts a")
	if criteria.Address != "" {
		chain = chain.Where("a.address = ?", criteria.Address)
	}

	// get total count
	var totalCount int64
	err := chain.Select("a.id").Count(&totalCount).Error
	if err != nil {
		logger.Error("failed to get total count", "err", err)
	}

	// get user ids
	rows, err := chain.Select("DISTINCT(a.id) id").
		Offset(int(criteria.Offset)).
		Limit(int(criteria.Limit)).
		Order("a.id DESC").
		Rows()
	if err != nil {
		logger.Error("failed to read user ids", "err", err)
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

	// get users by ids
	var ret []*model.Account
	if len(ids) == 0 {
		return []*model.Account{}, totalCount, nil
	}
	err = db.WithContext(ctx).
		Where("accounts.id IN (?)", ids).
		Order("accounts.id DESC").
		Find(&ret).Error
	if err != nil {
		logger.Error("failed to find user by ids", "err", err)
		return nil, 0, err
	}
	fmt.Println(ret)

	return ret, totalCount, nil
}

func (a *accountDB) FindByAddress(ctx context.Context, address string) (*model.Account, error) {
	logger := logging.FromContext(ctx)
	db := database.FromContext(ctx, a.db)
	logger.Debugw("account.db.FindByAddress", "address", address)

	var acc model.Account
	if err := db.WithContext(ctx).Where("address = ?", address).First(&acc).Error; err != nil {
		logger.Error("account.db.FindByAddress failed to find", "err", err)
		if database.IsRecordNotFoundErr(err) {
			return nil, database.ErrNotFound
		}
		return nil, err
	}
	return &acc, nil
}

// NewAccountDB creates a new account db with given db
func NewAccountDB(db *gorm.DB) AccountDB {
	return &accountDB{
		db: db,
	}
}
