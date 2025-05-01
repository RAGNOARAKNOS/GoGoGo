package dbase

import (
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/internal"
	"gorm.io/gorm"
)

type Boardgame struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Genre       internal.GameGenre `gorm:"type:game_genre"`
	// refactor to use dbase fkey table for genres
	Complexity   int
	Learnability int
	MinPlayers   int
	MaxPlayers   int
	BestPlayers  int
	Playtime     int
	Setuptime    int
	Designer     string
	BGGURL       string
	BGGRating    float32
	PublisherID  int
	Publisher    Publisher
	ImageURL     string
	Mechanics    []Tag       `gorm:"many2many:boardgame_mechanics;"`
	BitsNBobs    []Tag       `gorm:"many2many:boardgame_bits"`
	IsExpansion  bool        `gorm:"default:false"`
	ParentID     *uint       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Parent       *Boardgame  `gorm:"foreignkey:ParentID"`
	Expansions   []Boardgame `gorm:"foreignkey:ParentID"`
}
