package db

import (
	"fmt"

	"github.com/usawyer/urlShortener/models"
	"gorm.io/gorm"
)

type Database interface {
	AddUrl(models.Urls) error
	GetUrl(string) (string, error)
	FindUrl(string, string) (models.Urls, bool)
}

type pgClient struct {
	db *gorm.DB
}

func (p *pgClient) AddUrl(url models.Urls) error {
	res := p.db.Create(&url)
	return res.Error
}

func (p *pgClient) GetUrl(alias string) (string, error) {
	var url models.Urls
	res := p.db.Where("alias = ?", alias).First(&url)
	return url.Url, res.Error
}

func (p *pgClient) FindUrl(str string, columnName string) (models.Urls, bool) {
	var url models.Urls
	condition := fmt.Sprintf("%s = ?", columnName)
	res := p.db.Where(condition, str).Find(&url)
	if res.RowsAffected == 1 {
		return url, true
	}
	return url, false
}
