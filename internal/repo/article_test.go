package repo

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my-template-with-go/internal/entity"
	"regexp"
	"testing"
	"time"
)

type MockDBProvider struct {
	DBMain  *gorm.DB
	DBSlave *gorm.DB
}

func (m *MockDBProvider) GetDBMain() *gorm.DB {
	return m.DBMain
}

func (m *MockDBProvider) GetDBSlave() *gorm.DB {
	return m.DBSlave
}

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestArticleRepo_List(t *testing.T) {
	dbMain, _, err := setupMockDB()
	assert.NoError(t, err)

	dbSlave, mockSlave, err := setupMockDB()
	assert.NoError(t, err)

	// Set up the mock provider
	mockDBProvider := &MockDBProvider{
		DBMain:  dbMain,
		DBSlave: dbSlave,
	}

	repo := NewArticleRepo(mockDBProvider)
	rows := sqlmock.
		NewRows([]string{"created_at", "updated_at", "author", "title", "id"}).
		AddRow(time.Now().Unix(), time.Now().Unix(), "author 1", "title 1", 1).
		AddRow(time.Now().Unix(), time.Now().Unix(), "author 2", "title 2", 2)

	const sqlQuery = `
					SELECT * FROM "articles"
					`

	mockSlave.ExpectQuery(regexp.QuoteMeta(sqlQuery)).WillReturnRows(rows)

	// Act
	articles, err := repo.List()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, articles, 2)
	assert.NoError(t, mockSlave.ExpectationsWereMet())

}

func TestArticleRepo_Detail(t *testing.T) {
	dbMain, _, err := setupMockDB()
	assert.NoError(t, err)

	dbSlave, mockSlave, err := setupMockDB()
	assert.NoError(t, err)

	// Set up the mock provider
	mockDBProvider := &MockDBProvider{
		DBMain:  dbMain,
		DBSlave: dbSlave,
	}

	repo := NewArticleRepo(mockDBProvider)
	articleTest := &entity.Article{
		ID:        1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Author:    "author 1",
		Title:     "title 1",
	}

	rows := sqlmock.
		NewRows([]string{"created_at", "updated_at", "author", "title", "id"}).
		AddRow(articleTest.CreatedAt, articleTest.UpdatedAt, articleTest.Author, articleTest.Title, articleTest.ID)

	const sqlQuery = `
					SELECT * FROM "articles" WHERE id = $1 
					ORDER BY "articles"."id"
					LIMIT $2
					`

	mockSlave.ExpectQuery(regexp.QuoteMeta(sqlQuery)).
		WithArgs(articleTest.ID, 1).
		WillReturnRows(rows)

	// Act
	article, err := repo.Detail(articleTest.ID)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, article)
	assert.NoError(t, mockSlave.ExpectationsWereMet())

}

func TestArticleRepo_Create(t *testing.T) {
	dbMain, mockMain, err := setupMockDB()
	assert.NoError(t, err)

	dbSlave, _, err := setupMockDB()
	assert.NoError(t, err)

	// Set up the mock provider
	mockDBProvider := &MockDBProvider{
		DBMain:  dbMain,
		DBSlave: dbSlave,
	}

	repo := NewArticleRepo(mockDBProvider)
	newArticle := &entity.Article{
		ID:        1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Author:    "author",
		Title:     "This is a new article.",
	}

	const sqlInsert = `
					INSERT INTO "articles" ("created_at","updated_at","author","title","id") 
						VALUES ($1,$2,$3,$4,$5) RETURNING "id"
					`

	// https://github.com/DATA-DOG/go-sqlmock/issues/118
	mockMain.ExpectBegin()
	mockMain.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
		WithArgs(newArticle.CreatedAt, newArticle.UpdatedAt, newArticle.Author, newArticle.Title, newArticle.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newArticle.ID))
	mockMain.ExpectCommit()

	// Act
	err = repo.Create(newArticle)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, mockMain.ExpectationsWereMet())

}

func TestArticleRepo_Edit(t *testing.T) {
	dbMain, mockMain, err := setupMockDB()
	assert.NoError(t, err)

	dbSlave, _, err := setupMockDB()
	assert.NoError(t, err)

	// Set up the mock provider
	mockDBProvider := &MockDBProvider{
		DBMain:  dbMain,
		DBSlave: dbSlave,
	}

	repo := NewArticleRepo(mockDBProvider)
	editArticle := &entity.Article{
		ID:        1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Author:    "author",
		Title:     "This is a new article.",
	}

	newArticle := &entity.Article{
		ID:        1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Author:    "author",
		Title:     "This is a new article.",
	}

	const sqlInsert = `
					INSERT INTO "articles" ("created_at","updated_at","author","title","id") 
						VALUES ($1,$2,$3,$4,$5) RETURNING "id"
					`

	// https://github.com/DATA-DOG/go-sqlmock/issues/118
	mockMain.ExpectBegin()
	mockMain.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
		WithArgs(newArticle.CreatedAt, newArticle.UpdatedAt, newArticle.Author, newArticle.Title, newArticle.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newArticle.ID))
	mockMain.ExpectCommit()

	// Act
	err = repo.Create(newArticle)

	const sqlEdit = `
					UPDATE "articles" SET "author"=$1,"created_at"=$2,"updated_at"=$3 WHERE id = $4
					`

	mockMain.ExpectBegin()
	mockMain.ExpectExec(regexp.QuoteMeta(sqlEdit)).
		WithArgs(editArticle.Author, editArticle.CreatedAt, editArticle.UpdatedAt, editArticle.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mockMain.ExpectCommit()

	updateItems := map[string]interface{}{
		"author":     editArticle.Author,
		"created_at": editArticle.CreatedAt,
		"updated_at": editArticle.UpdatedAt,
	}
	// Act
	err = repo.Update(editArticle.ID, updateItems)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, mockMain.ExpectationsWereMet())

}
