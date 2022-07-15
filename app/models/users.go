package models

type Users struct {
	ID           uint   `json:"id" gorm:"primary_key;autoIncrement;uniqueIndex"`
	Email        string `gorm:"unique;not null"`
	Username     string
	Password     string
	Ip           string
	Balance      int64
	RegisterTime string
}
