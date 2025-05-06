package dbase

import "gorm.io/gorm"

type Gamecopy struct {
	gorm.Model
	BoardGameID uint       `gorm:"not null;index"`
	Game        *Boardgame `gorm:"foreignKey:BoardGameID"`
	Condition   int
	Love        int
	Notes       string `gorm:"type:text"`
}
