package dbase

import (
	"fmt"
	"log"

	"github.com/ragnoaraknos/GoGoGo/GoGameLibrary/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Query PSQL to see whether a custom data type exists
// return true if present
func enumExists(db *gorm.DB, enumName string) (bool, error) {
	var exists bool
	err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = $1)", enumName).Scan(&exists).Error
	if err != nil {
		return false, fmt.Errorf("ERROR checking if enum '%s' exists: %w", enumName, err)
	}
	return exists, nil
}

func ConnectDatabase() {
	databaseConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		internal.CFG.DBHost,
		internal.CFG.DBUser,
		internal.CFG.DBPassword,
		internal.CFG.DBName,
		internal.CFG.DBPort)

	db, err := gorm.Open(postgres.Open(databaseConnection), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	enumGameGenreExists, err := enumExists(db, "game_genre")
	if err != nil {
		fmt.Println(err)
		return
	}

	if !enumGameGenreExists {
		err = db.Exec(`
			CREATE TYPE game_genre as ENUM (
				'AbstractStrategy',
				'Cooperative',
				'DeckBuilder',
				'EconomicEngine',
				'Party',
				'RollAndMove',
				'SocialDeduction',
				'Thematic',
				'Wargame',
				'WorkerPlacement',
				'AreaControl',
				'Puzzle',
				'Dexterity',
				'Legacy');
		`).Error
		if err != nil {
			fmt.Println("ERROR creating ENUM types:", err)
			return
		}
	}

	DB = db

	// Use automigrate to create the 'games' tables - if one doesn't already exist
	DB.AutoMigrate(&Game{}, &Publisher{}, &Tag{}, &Boardgame{})
}
