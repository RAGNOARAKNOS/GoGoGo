package dbase

import "gorm.io/gorm"

type Mechanic struct {
	gorm.Model
	Name string
}
