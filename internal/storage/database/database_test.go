package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usawyer/urlShortener/models"
)

func TestPgClient_AddUrl(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %s", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		t.Fatalf("error opening gorm database: %s", err)
	}

	p := &pgClient{db: gdb}

	testUrl := models.Urls{
		Alias: "test",
		Url:   "https://example.com",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "urls"`).WithArgs(testUrl.Alias, testUrl.Url).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = p.AddUrl(testUrl)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPgClient_GetUrl(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %s", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		t.Fatalf("error opening gorm database: %s", err)
	}

	p := &pgClient{db: gdb}

	testUrl := models.Urls{
		Alias: "test",
		Url:   "https://example.com",
	}

	rows := sqlmock.NewRows([]string{"url"}).AddRow(testUrl.Url)
	expectedSQL := "SELECT (.+) FROM \"urls\" WHERE alias =(.+)"

	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)
	actualUrl, err := p.GetUrl(testUrl.Alias)

	assert.NoError(t, err)
	assert.Equal(t, testUrl.Url, actualUrl)
}

func TestPgClient_FindUrl(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %s", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		t.Fatalf("error opening gorm database: %s", err)
	}

	p := &pgClient{db: gdb}

	testUrl := models.Urls{
		Alias: "test",
		Url:   "https://example.com",
	}

	columnName := "url"
	rows := sqlmock.NewRows([]string{"alias", "url"}).AddRow(testUrl.Alias, testUrl.Url)
	expectedSQL := "SELECT (.+) FROM \"urls\" WHERE url =(.+)"

	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)
	actualUrl, ok := p.FindUrl(testUrl.Url, columnName)

	assert.True(t, ok)
	assert.Equal(t, testUrl, actualUrl)
}
