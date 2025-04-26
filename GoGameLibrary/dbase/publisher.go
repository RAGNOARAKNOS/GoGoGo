package dbase

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name string
	//Games []Boardgame `gorm:"foreignKey:PublisherID"`
}
