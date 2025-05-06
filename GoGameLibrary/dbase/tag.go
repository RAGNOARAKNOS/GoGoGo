package dbase

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex;not null"`
	Notes string `gorm:"type:text"`
}
