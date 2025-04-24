package internal

import (
	"database/sql/driver"
	"fmt"

	"slices"

	"github.com/go-playground/validator/v10"
)

type GameGenre string

const (
	GenreAreaControl      GameGenre = "AreaControl"
	GenreAbstractStrategy GameGenre = "AbstractStrategy"
	GenreCooperative      GameGenre = "Cooperative"
	GenreDeckBuild        GameGenre = "DeckBuilder"
	GenreDexterity        GameGenre = "Dexterity"
	GenreEconomic         GameGenre = "EconomicEngine"
	GenreLegacy           GameGenre = "Legacy"
	GenreParty            GameGenre = "Party"
	GenrePuzzle           GameGenre = "Puzzle"
	GenreRollMove         GameGenre = "RollAndMove"
	GenreSocialDeduction  GameGenre = "SocialDeduction"
	GenreThematic         GameGenre = "Thematic"
	GenreWargame          GameGenre = "Wargame"
	GenreWorkerPlacement  GameGenre = "WorkerPlacement"
)

// Scan method needed for reliable GORM database type conversion
func (g *GameGenre) Scan(value interface{}) error {
	if value == nil {
		*g = ""
		return nil
	}

	strVal, ok := value.(string)
	if !ok {
		return fmt.Errorf("FAILED to scan GameGenre: value must be a string")
	}

	*g = GameGenre(strVal)
	return nil
}

// Value method needed for reliable GORM database type conversion
func (g GameGenre) Value() (driver.Value, error) {
	return string(g), nil
}

// validGameGenres is used with the 'Validator' golib to ensure json is correctly structured
var validGameGenres = []string{
	string(GenreAreaControl),
	string(GenreAbstractStrategy),
	string(GenreCooperative),
	string(GenreDeckBuild),
	string(GenreDexterity),
	string(GenreEconomic),
	string(GenreLegacy),
	string(GenreParty),
	string(GenrePuzzle),
	string(GenreRollMove),
	string(GenreSocialDeduction),
	string(GenreThematic),
	string(GenreWargame),
	string(GenreWorkerPlacement),
}

// Used by Validator/v10 library to validate json request structure
func IsValidGameGenre(fl validator.FieldLevel) bool {
	genre := fl.Field().String()
	return slices.Contains(validGameGenres, genre)
}
