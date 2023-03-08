package database

import (
	accountDB "rena-server-v2/internal/account/database"
	accountModel "rena-server-v2/internal/account/model"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/registry/model"
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
	db        RegistryDB
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

func (s *DBSuite) TestSaveRegistry() {
	// given
	s.NoError(s.originDB.Create(&model.Tag{Name: "tag1"}).Error)
	registry := newRegistry("title1", "title1", "body", dUser, []string{"tag1", "tag2"})

	// when
	now := time.Now()
	err := s.db.SaveRegistry(nil, registry)

	// then
	s.NoError(err)
	find, err := s.db.FindRegistryBySlug(nil, registry.Slug)
	s.NoError(err)
	s.NotEqual(0, find.ID)
	s.Equal(registry.Slug, find.Slug)
	s.Equal(registry.Title, find.Title)
	s.Equal(registry.Body, find.Body)
	s.WithinDuration(now, find.CreatedAt, time.Second)
	s.WithinDuration(now, find.UpdatedAt, time.Second)
	s.Equal(int64(0), find.DeletedAtUnix)
	s.Equal(registry.Author, dUser)
	s.assertRegistryTag(find, []string{"tag1", "tag2"})
}

func (s *DBSuite) TestSaveRegistry_WithSameSlugAfterDeleted() {
	// given
	registry := newRegistry("title1", "title1", "body", dUser, []string{"tag1", "tag2"})
	s.NoError(s.db.SaveRegistry(nil, registry))
	s.NoError(s.db.DeleteRegistryBySlug(nil, dUser.ID, registry.Slug))

	registry2 := newRegistry(registry.Slug, registry.Title, registry.Body, dUser, []string{})

	// when
	err := s.db.SaveRegistry(nil, registry2)

	// then
	s.NoError(err)
}

func (s *DBSuite) TestSaveRegistry_FailIfDuplicateSlug() {
	// given
	registry := newRegistry("title1", "title1", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry))

	// when
	registry2 := newRegistry(registry.Slug, registry.Title, registry.Body, dUser, nil)
	err := s.db.SaveRegistry(nil, registry2)

	// then
	s.Error(err)
	s.Equal(database.ErrKeyConflict, err)
}

func (s *DBSuite) TestFindRegistryBySlug() {
	// given
	registry := newRegistry("title1", "title1", "body", dUser, []string{"tag1"})
	//now := time.Now()
	s.NoError(s.db.SaveRegistry(nil, registry))

	// when
	find, err := s.db.FindRegistryBySlug(nil, registry.Slug)

	// then
	s.NoError(err)
	s.assertRegistry(registry, find)
}

func (s *DBSuite) TestFindRegistryBySlug_FailIfNotExist() {
	// when
	find, err := s.db.FindRegistryBySlug(nil, "not-exist-slug")

	// then
	s.Nil(find)
	s.Error(err)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestFindRegistryBySlug_FailIfDeleted() {
	// given
	registry := newRegistry("title1", "title1", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry))
	_, err := s.db.FindRegistryBySlug(nil, registry.Slug)
	s.NoError(err)
	s.NoError(s.db.DeleteRegistryBySlug(nil, dUser.ID, registry.Slug))

	// when
	find, err := s.db.FindRegistryBySlug(nil, registry.Slug)

	// then
	s.Nil(find)
	s.Error(err)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestFindRegistries() {
	// given
	// User1
	// registry1 - tag1, tag2 <- second itr [0]
	// registry2 - tag1		 <- first itr [1]
	// registry3 - tag4
	// registry4 - tag3
	// registry5 - tag1		 <- first itr [0]
	// registry6 - tag1 (deleted)
	// User2
	// registry7 - tag1
	user1 := accountModel.Account{Username: "test-user1", Email: "test-user1@gmail.com", Password: "password"}
	s.NoError(s.accountDB.Save(nil, &user1))
	registry1 := newRegistry("registry1", "registry1", "body1", user1, []string{"tag1", "tag2"})
	s.NoError(s.db.SaveRegistry(nil, registry1))
	registry2 := newRegistry("registry2", "registry2", "body2", user1, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry2))
	registry3 := newRegistry("registry3", "registry3", "body3", user1, []string{"tag4"})
	s.NoError(s.db.SaveRegistry(nil, registry3))
	registry4 := newRegistry("registry4", "registry4", "body4", user1, []string{"tag3"})
	s.NoError(s.db.SaveRegistry(nil, registry4))
	registry5 := newRegistry("registry5", "registry5", "body5", user1, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry5))
	registry6 := newRegistry("registry6", "registry6", "body6", user1, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry6))
	s.NoError(s.db.DeleteRegistryBySlug(nil, user1.ID, registry6.Slug))

	user2 := accountModel.Account{Username: "test-user2", Email: "test-user2@gmail.com", Password: "password"}
	s.NoError(s.accountDB.Save(nil, &user2))
	registry7 := newRegistry("registry7", "registry7", "body7", user2, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry7))

	criteria := IterateRegistryCriteria{
		Tags:   []string{"tag1", "tag2"},
		Author: user1.Username,
		Offset: 0,
		Limit:  2,
	}

	// when : first iteration
	results, total, err := s.db.FindRegistries(nil, criteria)

	// then
	s.NoError(err)
	s.Equal(int64(3), total)
	s.Equal(2, len(results))
	s.assertRegistry(registry5, results[0])
	s.assertRegistry(registry2, results[1])

	// second iteration
	criteria.Offset = criteria.Offset + uint(len(results))
	results, total, err = s.db.FindRegistries(nil, criteria)

	// then
	s.NoError(err)
	s.Equal(int64(3), total)
	s.Equal(1, len(results))
	s.assertRegistry(registry1, results[0])
}

func (s *DBSuite) TestDeleteRegistryBySlug() {
	// given
	registry := newRegistry("title1", "title1", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry))

	// when
	err := s.db.DeleteRegistryBySlug(nil, dUser.ID, registry.Slug)

	// then
	s.NoError(err)
	find, err := s.db.FindRegistryBySlug(nil, registry.Slug)
	s.Nil(find)
	s.Equal(database.ErrNotFound, err)
}

func (s *DBSuite) TestDeleteRegistryBySlug_FailIfNotExist() {
	// given
	registry := newRegistry("title1", "title1", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry))
	registry2 := newRegistry("title2", "title2", "body", dUser, []string{"tag1"})
	s.NoError(s.db.SaveRegistry(nil, registry2))
	s.NoError(s.db.DeleteRegistryBySlug(nil, registry2.Author.ID, registry2.Slug))

	cases := []struct {
		AuthorID uint
		Slug     string
	}{
		{
			AuthorID: dUser.ID,
			Slug:     "not-exist-slug",
		}, {
			AuthorID: dUser.ID + 1000,
			Slug:     registry.Slug,
		}, {
			AuthorID: dUser.ID,
			Slug:     registry2.Slug,
		},
	}

	for _, tc := range cases {
		// when
		err := s.db.DeleteRegistryBySlug(nil, tc.AuthorID, tc.Slug)

		// then
		s.Error(err)
		s.Equal(database.ErrNotFound, err)
	}
}

func (s *DBSuite) assertRegistry(expected, actual *model.Registry) {
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
	s.assertRegistryTag(actual, tags)
}

func (s *DBSuite) assertRegistryTag(registry *model.Registry, tags []string) {
	s.Equal(len(registry.Tags), len(tags))
	if len(registry.Tags) == 0 {
		return
	}
	m := make(map[string]struct{})
	for _, tag := range registry.Tags {
		m[tag.Name] = struct{}{}
	}
	for _, tag := range tags {
		_, ok := m[tag]
		s.True(ok)
	}
}

func newRegistry(slug, title, body string, author accountModel.Account, tags []string) *model.Item {
	var tagArr []*model.Tag
	for _, tag := range tags {
		tagArr = append(tagArr, &model.Tag{Name: tag})
	}
	return &model.Registry{
		Slug:   slug,
		Title:  title,
		Body:   body,
		Author: author,
		Tags:   tagArr,
	}
}
