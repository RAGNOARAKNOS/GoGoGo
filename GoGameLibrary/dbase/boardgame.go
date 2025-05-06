package dbase

import (
	"gorm.io/gorm"
)

type Boardgame struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	ImageURL    string
	// Playability
	Complexity   int
	Learnability int
	Playtime     int
	Setuptime    int
	// Players
	MinPlayers  int
	MaxPlayers  int
	BestPlayers int
	// Provenance
	Designer  string
	BGGURL    string
	BGGRating float32
	// Relationship: A Boardgame (optionally) belongs to one Publisher (1-1)
	PublisherID *uint      `gorm:"not null; index"`        // FK to Publisher (nullable/optional)
	Publisher   *Publisher `gorm:"foreignKey:PublisherID"` // 1-1 to Publisher (nullable optional)
	// Relationship: Boardgames (optionally) belongs to many Genres (n-n)
	Genres []*Tag `gorm:"many2many:boardgame_genre;"` // Many-to-many with Tag for Genres
	//Genre       internal.GameGenre `gorm:"type:game_genre"`
	// Relationship: Boardgames (optionally) belongs to many Tag(Mechanics) (n-n)
	Mechanics []*Tag `gorm:"many2many:boardgame_mechanic;"`
	// Relationship: Boardgames (optionally) belongs to many Tag(Bits)
	Bits []*Tag `gorm:"many2many:boardgame_bit"`
	// Relationship: A Boardgame (optionally) belongs to many Boardgame(Expansions) (1-n)
	BasegameID *uint        `gorm:"index"` // Pointer allows null for 'base' games (non-expansions)
	Expansions []*Boardgame `gorm:"foreignkey:BasegameID"`
}
