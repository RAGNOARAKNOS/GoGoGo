package dtos

import "github.com/ragnoaraknos/GoGoGo/GoGameLibrary/internal"

type NewBoardgameRequest struct {
	Title       string             `json:"title" validate:"required"`
	Description string             `json:"description"`
	Genre       internal.GameGenre `json:"genre" validate:"required"`
	Complexity  int                `json:"complexity" validate:"gte=0,lte=10"`
	MinPlayers  int                `json:"players_min" validate:"gte=1,lte=99"`
	MaxPlayers  int                `json:"players_max" validate:"gte=1,lte=99"`
	BestPlayers int                `json:"players_best" validate:"gte=1,lte=99"`
	Playtime    int                `json:"playtime"`
	Designer    string             `json:"designer"`
	PublisherID int                `json:"publisher_id"`
	ImageURL    string             `json:"image_url" validate:"url"`
	// ParentID    *uint               `json:"parentid"`
	//BoardGameGeekURL    string             `json:"bgg_url" validate:"url"`  //IDEA! get from BGG API
	//BoardGameGeekRating float32            `json:"bgg_rating" validate:"gte=0,lte=10"`
}

type UpdateBoardgameRequest struct {
}

type BoardgameResponse struct {
	ID           uint               `json:"id"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	Genre        internal.GameGenre `json:"genre"`
	Complexity   int                `json:"complexity"`
	Learnability int                `json:"learnability"`
	MinPlayers   int                `json:"players_min"`
	MaxPlayers   int                `json:"players_max"`
	BestPlayers  int                `json:"players_best"`
	Playtime     int                `json:"playtime"`
	Designer     string             `json:"designer"`
	PublisherID  int                `json:"publisher_id"`
	ImageURL     string             `json:"image_url"`
	ParentID     int                `json:"parentid"`
}
