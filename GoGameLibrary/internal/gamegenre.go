package internal

import (
	"database/sql/driver"
	"fmt"
)

type GameGenre string

const (
	GenreAbstractStrategy GameGenre = "AbstractStrategy"
	GenreCooperative      GameGenre = "Cooperative"
	GenreDeckBuild        GameGenre = "DeckBuilder"
	GenreEconomic         GameGenre = "EconomicEngine"
	GenreParty            GameGenre = "Party"
	GenreRollMove         GameGenre = "RollAndMove"
	GenreSocialDeduction  GameGenre = "SocialDeduction"
	GenreThematic         GameGenre = "Thematic"
	GenreWargame          GameGenre = "Wargame"
	GenreWorkerPlacement  GameGenre = "WorkerPlacement"
	GenreAreaControl      GameGenre = "AreaControl"
	GenrePuzzle           GameGenre = "Puzzle"
	GenreDexterity        GameGenre = "Dexterity"
	GenreLegacy           GameGenre = "Legacy"
)

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

func (g GameGenre) Value() (driver.Value, error) {
	return string(g), nil
}
