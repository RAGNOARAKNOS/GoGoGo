package dbase

import "gorm.io/gorm"

type Gamecopy struct {
	gorm.Model
	Game      Boardgame
	Condition *int
	Love      *int
}
