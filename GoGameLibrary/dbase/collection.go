package dbase

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name  string     `gorm:"not null"`
	Games []Gamecopy `gorm:"foreignKey:GamecopyID"`
}
