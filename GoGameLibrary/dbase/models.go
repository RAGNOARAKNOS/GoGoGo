package dbase

import (
	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/internal"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Title     string `json:"title"`
	Studio    string `json:"studio"`
	Publisher string `json:"publisher"`
}

type Publisher struct {
	gorm.Model
	Name  string      `json:"name"`
	Games []Boardgame `json:"games,omitempty"`
}

type Boardgame struct {
	gorm.Model
	Title       string             `json:"title" gorm:"not null"`
	Description string             `json:"description"`
	Genre       internal.GameGenre `gorm:"type:game_genre" json:"genre"`
	Complexity  int                `json:"complexity"`
	MinPlayers  int                `json:"players_min"`
	MaxPlayers  int                `json:"players_max"`
	BestPlayers int                `json:"players_best"`
	Playtime    int                `json:"playtime"`
	Designer    string             `json:"designer"`
	BGGURL      string             `json:"bgg_url"`
	BGGRating   float32            `json:"bgg_rating"`
	PublisherID int                `json:"publisher_id"`
	Publisher   Publisher
	ImageURL    string      `json:"image_url"`
	Mechanics   []Mechanic  `gorm:"many2many:boardgame_mechanics;"`
	IsExpansion bool        `gorm:"default:false"`
	ParentID    *uint       `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Parent      *Boardgame  `gorm:"foreignkey:ParentID"`
	Expansions  []Boardgame `gorm:"foreignkey:ParentID"`
}
