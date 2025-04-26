package dbase

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Title     string `json:"title"`
	Studio    string `json:"studio"`
	Publisher string `json:"publisher"`
}
