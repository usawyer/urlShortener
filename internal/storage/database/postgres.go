package db

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/usawyer/urlShortener/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func New(zapLogger *zap.Logger) Database {
	zapLogger = zapLogger.Named("PostgreSQL")
	dsn := makeDsnStr()
	gormLogger := zapgorm2.New(zapLogger)

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
		return &PgClient{db: db}
	}
	zapLogger.Fatal("Error open db")
	return nil
}

func makeDsnStr() string {
	parameters := map[string]string{
		"host":     getEnv("DB_HOST", "localhost"),
		"user":     getEnv("POSTGRES_USER", "postgres"),
		"password": getEnv("POSTGRES_PASSWORD", "postgres"),
		"dbname":   getEnv("POSTGRES_DB", "test"),
		"port":     getEnv("DB_PORT", "5432"),
		"sslmode":  "disable",
		"TimeZone": "Asia/Novosibirsk",
	}

	var pairs []string
	for key, value := range parameters {
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(pairs, " ")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
