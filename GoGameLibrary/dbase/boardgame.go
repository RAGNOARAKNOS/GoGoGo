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
	Complexity  int
	MinPlayers  int
	MaxPlayers  int
	BestPlayers int
	Playtime    int
	Designer    string
	BGGURL      string
	BGGRating   float32
	PublisherID int
	Publisher   Publisher
	ImageURL    string
	Mechanics   []Mechanic  `gorm:"many2many:boardgame_mechanics;"`
	IsExpansion bool        `gorm:"default:false"`
	ParentID    *uint       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Parent      *Boardgame  `gorm:"foreignkey:ParentID"`
	Expansions  []Boardgame `gorm:"foreignkey:ParentID"`
}
