package database

import (
	accountDB "rena-server-v2/internal/account/database"
	accountModel "rena-server-v2/internal/account/model"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/item/model"
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
	db        ItemDB
	accountDB accountDB.AccountDB
	originDB  *gorm.DB
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(DBSuite))
}

func (s *DBSuite) SetupSuite() {
	logging.SetLevel(zapcore.FatalLevel)
	s.originDB = database.NewTestDatabase(s.T(), true)
	s.db = &itemDB{db: s.originDB}
	s.accountDB = accountDB.NewAccountDB(s.originDB)
}

func (s *DBSuite) SetupTest() {
	s.NoError(database.DeleteRecordAll(s.T(), s.originDB, []string{
		"item_tags", "item_id > 0",
		"tags", "id > 0",
		"items", "id > 0",
		"accounts", "id > 0",
	}))
	s.NoError(s.accountDB.Save(nil, &dUser))
}

func (s *DBSuite) TestSaveItem() {
	// given
	s.NoError(s.originDB.Create(&model.Tag{Name: "tag1"}).Error)
	item := newItem("title1", "title1", "body", dUser, []string{"tag1", "tag2"})

	// when
	now := time.Now()
	err := s.db.SaveItem(nil, item)

	// then
	s.NoError(err)
	find, err := s.db.FindItemBySlug(nil, item.Slug)
	s.NoError(err)
	s.NotEqual(0, find.ID)
	s.Equal(item.Slug, find.Slug)
	s.Equal(item.Title, find.Title)
	s.Equal(item.Body, find.Body)
	s.WithinDuration(now, find.CreatedAt, time.Second)
	s.WithinDuration(now, find.UpdatedAt, time.Second)
	s.Equal(int64(0), find.DeletedAtUnix)
	s.Equal(item.Author, dUser)
	s.assertItemTag(find, []string{"tag1", "tag2"})
}

func (s *DBSuite) TestSaveItem_WithSameSlugAfterDeleted() {
	// given
	item := newItem("title1", "title1", "body", dUser, []string{"tag1", "tag2"})
	s.NoError(s.db.SaveItem(nil, item))
	s.NoError(s.db.DeleteItemBySlug(nil, dUser.ID, item.Slug))

	item2 := newItem(item.Slug, item.Title, item.Body, dUser, []string{})

	// when
	err := s.db.SaveItem(nil, item2)

	// then
	s.NoError(err)
}

func (s *DBSuite) TestSaveItem_FailIfDuplicateSlug() {
	// given
	item := newItem("title1", "title1", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item))

	// when
	item2 := newItem(item.Slug, item.Title, item.Body, dUser, nil)
	err := s.db.SaveItem(nil, item2)

	// then
	s.Error(err)
	s.Equal(database.ErrKeyConflict, err)
}

func (s *DBSuite) TestFindItemBySlug() {
	// given
	item := newItem("title1", "title1", "body", dUser, []string{"tag1"})
	//now := time.Now()
	s.NoError(s.db.SaveItem(nil, item))

	// when
	find, err := s.db.FindItemBySlug(nil, item.Slug)

	// then
	s.NoError(err)
	s.assertItem(item, find)
}

func (s *DBSuite) TestFindItemBySlug_FailIfNotExist() {
	// when
	find, err := s.db.FindItemBySlug(nil, "not-exist-slug")

	// then
	s.Nil(find)
	s.Error(err)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestFindItemBySlug_FailIfDeleted() {
	// given
	item := newItem("title1", "title1", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item))
	_, err := s.db.FindItemBySlug(nil, item.Slug)
	s.NoError(err)
	s.NoError(s.db.DeleteItemBySlug(nil, dUser.ID, item.Slug))

	// when
	find, err := s.db.FindItemBySlug(nil, item.Slug)

	// then
	s.Nil(find)
	s.Error(err)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestFindItems() {
	// given
	// User1
	// item1 - tag1, tag2 <- second itr [0]
	// item2 - tag1		 <- first itr [1]
	// item3 - tag4
	// item4 - tag3
	// item5 - tag1		 <- first itr [0]
	// item6 - tag1 (deleted)
	// User2
	// item7 - tag1
	user1 := accountModel.Account{Username: "test-user1", Email: "test-user1@gmail.com", Password: "password"}
	s.NoError(s.accountDB.Save(nil, &user1))
	item1 := newItem("item1", "item1", "body1", user1, []string{"tag1", "tag2"})
	s.NoError(s.db.SaveItem(nil, item1))
	item2 := newItem("item2", "item2", "body2", user1, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item2))
	item3 := newItem("item3", "item3", "body3", user1, []string{"tag4"})
	s.NoError(s.db.SaveItem(nil, item3))
	item4 := newItem("item4", "item4", "body4", user1, []string{"tag3"})
	s.NoError(s.db.SaveItem(nil, item4))
	item5 := newItem("item5", "item5", "body5", user1, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item5))
	item6 := newItem("item6", "item6", "body6", user1, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item6))
	s.NoError(s.db.DeleteItemBySlug(nil, user1.ID, item6.Slug))

	user2 := accountModel.Account{Username: "test-user2", Email: "test-user2@gmail.com", Password: "password"}
	s.NoError(s.accountDB.Save(nil, &user2))
	item7 := newItem("item7", "item7", "body7", user2, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item7))

	criteria := IterateItemCriteria{
		Tags:   []string{"tag1", "tag2"},
		Author: user1.Username,
		Offset: 0,
		Limit:  2,
	}

	// when : first iteration
	results, total, err := s.db.FindItems(nil, criteria)

	// then
	s.NoError(err)
	s.Equal(int64(3), total)
	s.Equal(2, len(results))
	s.assertItem(item5, results[0])
	s.assertItem(item2, results[1])

	// second iteration
	criteria.Offset = criteria.Offset + uint(len(results))
	results, total, err = s.db.FindItems(nil, criteria)

	// then
	s.NoError(err)
	s.Equal(int64(3), total)
	s.Equal(1, len(results))
	s.assertItem(item1, results[0])
}

func (s *DBSuite) TestDeleteItemBySlug() {
	// given
	item := newItem("title1", "title1", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item))

	// when
	err := s.db.DeleteItemBySlug(nil, dUser.ID, item.Slug)

	// then
	s.NoError(err)
	find, err := s.db.FindItemBySlug(nil, item.Slug)
	s.Nil(find)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestDeleteItemBySlug_FailIfNotExist() {
	// given
	item := newItem("title1", "title1", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item))
	item2 := newItem("title2", "title2", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveItem(nil, item2))
	s.NoError(s.db.DeleteItemBySlug(nil, item2.Author.ID, item2.Slug))

	cases := []struct {
		AuthorID uint
		Slug     string
	}{
		{
			AuthorID: dUser.ID,
			Slug:     "not-exist-slug",
		}, {
			AuthorID: dUser.ID + 1000,
			Slug:     item.Slug,
		}, {
			AuthorID: dUser.ID,
			Slug:     item2.Slug,
		},
	}

	for _, tc := range cases {
		// when
		err := s.db.DeleteItemBySlug(nil, tc.AuthorID, tc.Slug)

		// then
		s.Error(err)
		s.Equal(database.ErrNotFound, err)
	}
}

func (s *DBSuite) assertItem(expected, actual *model.Item) {
	s.Equal(expected.Slug, actual.Slug)
	s.Equal(expected.Title, actual.Title)
	s.Equal(expected.Body, actual.Body)
	s.WithinDuration(expected.CreatedAt, actual.CreatedAt, time.Second)
	s.WithinDuration(expected.UpdatedAt, actual.UpdatedAt, time.Second)
	s.Equal(expected.DeletedAtUnix, actual.DeletedAtUnix)
	s.Equal(expected.Author.ID, actual.Author.ID)
	s.Equal(expected.Author.Email, actual.Author.Email)
	s.Equal(expected.Author.Username, actual.Author.Username)
	var tags []string
	for _, tag := range expected.Tags {
		tags = append(tags, tag.Name)
	}
	s.assertItemTag(actual, tags)
}

func (s *DBSuite) assertItemTag(item *model.Item, tags []string) {
	s.Equal(len(item.Tags), len(tags))
	if len(item.Tags) == 0 {
		return
	}
	m := make(map[string]struct{})
	for _, tag := range item.Tags {
		m[tag.Name] = struct{}{}
	}
	for _, tag := range tags {
		_, ok := m[tag]
		s.True(ok)
	}
}

func newItem(slug, title, body string, author accountModel.Account, tags []string) *model.Item {
	var tagArr []*model.Tag
	for _, tag := range tags {
		tagArr = append(tagArr, &model.Tag{Name: tag})
	}
	return &model.Item{
		Slug:   slug,
		Title:  title,
		Body:   body,
		Author: author,
		Tags:   tagArr,
	}
}
