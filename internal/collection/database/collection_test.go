package database

import (
	accountDB "rena-server-v2/internal/account/database"
	accountModel "rena-server-v2/internal/account/model"
	"rena-server-v2/internal/collection/model"
	"rena-server-v2/internal/database"
	"rena-server-v2/pkg/logging"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

var dUser = accountModel.Account{
	Username: "user1",
	Email:    "user1@gmail.com",
	Password: "password",
}

type DBSuite struct {
	suite.Suite
	db        CollectionDB
	accountDB accountDB.AccountDB
	originDB  *gorm.DB
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(DBSuite))
}

func (s *DBSuite) SetupSuite() {
	logging.SetLevel(zapcore.FatalLevel)
	s.originDB = database.NewTestDatabase(s.T(), true)
	s.db = &collectionDB{db: s.originDB}
	s.accountDB = accountDB.NewAccountDB(s.originDB)
}

func (s *DBSuite) SetupTest() {
	s.NoError(database.DeleteRecordAll(s.T(), s.originDB, []string{
		"collections", "id > 0",
		"comments", "id > 0",
		"article_tags", "article_id > 0",
		"tags", "id > 0",
		"articles", "id > 0",
		"accounts", "id > 0",
	}))
	s.NoError(s.accountDB.Save(nil, &dUser))
}

func (s *DBSuite) TestSaveCollection() {
	// given
	collection := newCollection("0x382...C", "name1", "name1", "description", "image")

	// when
	now := time.Now()
	err := s.db.SaveCollection(nil, collection)

	// then
	s.NoError(err)
	find, err := s.db.FindCollectionBySlug(nil, collection.Slug)
	s.NoError(err)
	s.NotEqual(0, find.ID)
	s.Equal(collection.Slug, find.Slug)
	s.Equal(collection.Name, find.Name)
	s.Equal(collection.Description, find.Description)
	s.Equal(collection.ImageURL, find.ImageURL)
	s.WithinDuration(now, find.CreatedAt, time.Second)
	s.WithinDuration(now, find.UpdatedAt, time.Second)
}

func (s *DBSuite) TestSaveCollection_WithSameSlugAfterDeleted() {
	// given
	collection := newCollection("0x382...C", "name1", "name1", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection))
	s.NoError(s.db.DeleteCollectionBySlug(nil, collection.Slug))

	collection2 := newCollection(collection.ContractAddress, collection.Slug, collection.Name, collection.Description, collection.ImageURL)

	// when
	err := s.db.SaveCollection(nil, collection2)

	// then
	s.NoError(err)
}

func (s *DBSuite) TestSaveCollection_FailIfDuplicateSlug() {
	// given
	collection := newCollection("0x382...C", "name1", "name1", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection))

	// when
	collection2 := newCollection("0x382...C", "name2", "name2", "description", "image")
	err := s.db.SaveCollection(nil, collection2)

	// then
	s.Error(err)
	s.Equal(database.ErrKeyConflict, err)
}

func (s *DBSuite) TestFindCollectionBySlug() {
	// given
	collection := newCollection("0x382...C", "name1", "name1", "description", "image")
	//now := time.Now()
	s.NoError(s.db.SaveCollection(nil, collection))

	// when
	find, err := s.db.FindCollectionBySlug(nil, collection.Slug)

	// then
	s.NoError(err)
	s.assertCollection(collection, find)
}

func (s *DBSuite) TestFindCollectionBySlug_FailIfNotExist() {
	// when
	find, err := s.db.FindCollectionBySlug(nil, "not-exist-slug")

	// then
	s.Nil(find)
	s.Error(err)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestFindCollectionBySlug_FailIfDeleted() {
	// given
	collection := newCollection("0x382...C", "name1", "name1", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection))
	_, err := s.db.FindCollectionBySlug(nil, collection.Slug)
	s.NoError(err)
	s.NoError(s.db.DeleteCollectionBySlug(nil, collection.Slug))

	// when
	find, err := s.db.FindCollectionBySlug(nil, collection.Slug)

	// then
	s.Nil(find)
	s.Error(err)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestFindCollections() {
	// given
	// collection1 - tag1, tag2 <- second itr [0]
	// collection2 - tag1		 <- first itr [1]
	// collection3 - tag4
	// collection4 - tag3
	// collection5 - tag1		 <- first itr [0]
	// collection6 - tag1 (deleted)
	// User2
	// collection7 - tag1
	collection1 := newCollection("0x382...C", "name1", "name1", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection1))
	collection2 := newCollection("0x382...C", "name2", "name2", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection2))
	collection3 := newCollection("0x382...C", "name3", "name3", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection3))
	collection4 := newCollection("0x382...C", "name4", "name4", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection4))
	collection5 := newCollection("0x382...C", "name5", "name5", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection5))
	collection6 := newCollection("0x382...C", "name6", "name6", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection6))
	s.NoError(s.db.DeleteCollectionBySlug(nil, collection6.Slug))

	user2 := accountModel.Account{Username: "test-user2", Email: "test-user2@gmail.com", Password: "password"}
	s.NoError(s.accountDB.Save(nil, &user2))
	collection7 := newCollection("0x382...C", "name7", "name7", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection7))

	criteria := IterateCollectionCriteria{
		Slug:   collection1.Slug,
		Offset: 0,
		Limit:  2,
	}

	// when : first iteration
	results, total, err := s.db.FindCollections(nil, criteria)

	// then
	s.NoError(err)
	s.Equal(int64(3), total)
	s.Equal(2, len(results))
	s.assertCollection(collection5, results[0])
	s.assertCollection(collection2, results[1])

	// second iteration
	criteria.Offset = criteria.Offset + uint(len(results))
	results, total, err = s.db.FindCollections(nil, criteria)

	// then
	s.NoError(err)
	s.Equal(int64(3), total)
	s.Equal(1, len(results))
	s.assertCollection(collection1, results[0])
}

func (s *DBSuite) TestDeleteCollectionBySlug() {
	// given
	collection := newCollection("0x382...C", "name1", "name1", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection))

	// when
	err := s.db.DeleteCollectionBySlug(nil, collection.Slug)

	// then
	s.NoError(err)
	find, err := s.db.FindCollectionBySlug(nil, collection.Slug)
	s.Nil(find)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestDeleteCollectionBySlug_FailIfNotExist() {
	// given
	collection := newCollection("0x382...C", "name1", "name1", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection))
	collection2 := newCollection("0x382...C", "name2", "name2", "description", "image")
	s.NoError(s.db.SaveCollection(nil, collection2))
	s.NoError(s.db.DeleteCollectionBySlug(nil, collection2.Slug))

	cases := []struct {
		Slug string
	}{
		{
			Slug: "not-exist-slug",
		}, {
			Slug: collection.Slug,
		}, {
			Slug: collection.Slug,
		},
	}

	for _, tc := range cases {
		// when
		err := s.db.DeleteCollectionBySlug(nil, tc.Slug)

		// then
		s.Error(err)
		s.Equal(database.ErrNotFound, err)
	}
}

func (s *DBSuite) assertCollection(expected, actual *model.Collection) {
	s.Equal(expected.ContractAddress, actual.ContractAddress)
	s.Equal(expected.Slug, actual.Slug)
	s.Equal(expected.Name, actual.Name)
	s.Equal(expected.Description, actual.Description)
	s.Equal(expected.ImageURL, actual.ImageURL)
	s.WithinDuration(expected.CreatedAt, actual.CreatedAt, time.Second)
	s.WithinDuration(expected.UpdatedAt, actual.UpdatedAt, time.Second)
}

func newCollection(contractAddress, slug, name, description string, imageURL string) *model.Collection {
	return &model.Collection{
		ContractAddress: contractAddress,
		Slug:            slug,
		Name:            name,
		Description:     description,
		ImageURL:        imageURL,
	}
}
