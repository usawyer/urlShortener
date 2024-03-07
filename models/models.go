package models

type Urls struct {
	alias string `gorm:"index:idx_alias;unique;not null"`
	url   string
}

type Request struct {
}

type Response struct {
	Alias string `json:"alias"`
}
