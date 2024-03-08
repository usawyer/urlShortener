package models

type Urls struct {
	Alias string `gorm:"index:idx_alias;unique;not null"`
	Url   string
}

type Response struct {
	Alias string `json:"alias"`
}
