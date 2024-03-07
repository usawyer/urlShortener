package db

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/usawyer/urlShortener/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"os"
	"time"
)

type Database interface {
	AddUrl(string) (string, error)
	GetUrl(string) (string, error)
}

type pgClient struct {
	db *gorm.DB
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func New(zapLogger *zap.Logger) *pgClient {
	connectionParams := map[string]string{
		"host":     getEnv("DB_HOST", "localhost"),
		"user":     getEnv("POSTGRES_USER", "postgres"),
		"password": getEnv("POSTGRES_PASSWORD", "postgres"),
		"dbname":   getEnv("POSTGRES_DB", "test"),
		"port":     getEnv("DB_PORT", "5432"),
		"sslmode":  "disable",
		"TimeZone": "Asia/Novosibirsk",
	}
	gormLogger := zapgorm2.New(zapLogger)
	var dsn string

	for key, value := range connectionParams {
		dsn += fmt.Sprintf("%s=%s ", key, value)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err != nil {
			zapLogger.Warn("Error open", zap.Error(err))
			continue
		}
		db = db.Debug()
		err = db.AutoMigrate(&models.Urls{})
		if err != nil {
			zapLogger.Error(err.Error())
		}
		zapLogger.Info("migrate ok")
		return &pgClient{db: db}
	}
	zapLogger.Fatal("Error open db")
	return nil
}

func (p *pgClient) AddUrl(url string) (string, error) {
	res := p.db.Create(&url)
	return "", res.Error
}

func (p *pgClient) GetUrl(alias string) (string, error) {
	var article models.Urls
	res := p.db.First(&article, alias)
	return "article", res.Error
}

func (p *pgClient) RemoveArticle(id int) error {
	res := p.db.Delete(&models.Urls{}, id)
	if res.Error == nil && res.RowsAffected != 1 {
		return errors.New("article with such id doesn't exist")

	}
	return res.Error
}
