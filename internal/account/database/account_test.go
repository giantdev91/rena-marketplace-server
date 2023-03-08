package database

import (
	"rena-server-v2/internal/account/model"
	"rena-server-v2/internal/database"
	"rena-server-v2/pkg/logging"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type DBSuite struct {
	suite.Suite
	db       AccountDB
	originDB *gorm.DB
}

func (s *DBSuite) SetupSuite() {
	logging.SetLevel(zapcore.FatalLevel)
	s.originDB = database.NewTestDatabase(s.T(), true)
	s.db = &accountDB{
		db: s.originDB,
	}
}

func (s *DBSuite) SetupTest() {
	s.originDB.Where("id > 0").Delete(&model.Account{})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(DBSuite))
}

func (s *DBSuite) TestSave() {
	// given
	acc := model.Account{
		Username: "user1",
		Email:    "user@gmail.com",
	}

	// when
	now := time.Now()
	err := s.db.Save(nil, &acc)

	// then
	s.NoError(err)
	find, err := s.db.FindByAddress(nil, acc.Email)
	s.NoError(err)
	s.Equal(acc.ID, find.ID)
	s.Equal(acc.Address, find.Address)
	s.Equal(acc.Username, find.Username)
	s.Equal(acc.Email, find.Email)
	s.Equal(acc.Bio, find.Bio)
	s.Equal(acc.Image, find.Image)
	s.WithinDuration(now, find.CreatedAt, time.Second)
	s.WithinDuration(now, find.UpdatedAt, time.Second)
}

func (s *DBSuite) TestSave_ErrorIfExistEmail() {
	// given
	email := "user1@email.com"
	err := s.db.Save(nil, &model.Account{
		Username: "user1",
		Email:    email,
	})
	s.NoError(err)

	// when
	err = s.db.Save(nil, &model.Account{
		Username: "user2",
		Email:    email,
	})

	// then
	s.Error(err)
	s.Equal(database.ErrKeyConflict, err)
}

func (s *DBSuite) TestUpdate() {
	// given
	acc := model.Account{
		Username: "user1",
		Email:    "user@gmail.com",
	}
	s.NoError(s.db.Save(nil, &acc))

	updated := &model.Account{
		Username: "updated-user1",
		Email:    "updated-email@gamil.com",
		Bio:      "updated-bio",
		Image:    "updated-image",
	}

	// when
	err := s.db.Update(nil, acc.Email, updated)

	// then
	s.NoError(err)

	find, err := s.db.FindByAddress(nil, acc.Email)
	s.NoError(err)
	// unchanged fields
	s.Equal(acc.Address, find.Address)
	s.Equal(acc.Email, find.Email)

	// updated fields
	s.Equal(updated.Username, find.Username)
	s.Equal(updated.Bio, find.Bio)
	s.Equal(updated.Image, find.Image)
}

func (s *DBSuite) TestUpdate_FailIfNotExist() {
	// when
	err := s.db.Update(nil, "unknown@emai.com", &model.Account{
		Username: "updated",
	})

	// then
	s.Error(err)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestFindByAddress() {
	// given
	now := time.Now()
	acc := model.Account{
		Username: "user1",
		Email:    "user@gmail.com",
	}
	s.NoError(s.db.Save(nil, &acc))

	// when
	find, err := s.db.FindByAddress(nil, acc.Email)

	// then
	s.NoError(err)
	s.Equal(acc.Address, find.Address)
	s.Equal(acc.Username, find.Username)
	s.Equal(acc.Email, find.Email)
	s.Empty(find.Bio)
	s.Empty(find.Image)
	s.WithinDuration(now, find.CreatedAt, time.Second)
	s.WithinDuration(now, find.UpdatedAt, time.Second)
}

func (s *DBSuite) TestFindByAddress_ErrorIfNotExist() {
	// when
	find, err := s.db.FindByAddress(nil, "unknown@email.com")

	// then
	s.Nil(find)
	s.Error(err)
	s.Equal(database.ErrNotFound, err)
}
