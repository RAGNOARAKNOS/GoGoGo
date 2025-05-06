package dtos

type NewBoardgameRequest struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description"`
	ImageURL     string `json:"image_url" validate:"url"`
	Complexity   int    `json:"complexity" validate:"gte=0,lte=10"`
	Learnability int    `json:"learnability" validate:"gte=0,lte=10"`
	Playtime     int    `json:"playtime" validate:"gte=1"`
	Setuptime    int    `json:"setuptime"`
	MinPlayers   int    `json:"players_min" validate:"gte=1,lte=99"`
	MaxPlayers   int    `json:"players_max" validate:"gte=1,lte=99"`
	BestPlayers  int    `json:"players_best" validate:"gte=1,lte=99"`
	Designer     string `json:"designer"`
	//BoardGameGeekURL    string  `json:"bgg_url" validate:"url"`             //IDEA! get from BGG API
	//BoardGameGeekRating float32 `json:"bgg_rating" validate:"gte=0,lte=10"` //IDEA! get from BGG API
	PublisherID  *uint  `json:"publisher_id,omitempty" validate:"omitempty,gt=0"`
	GenreIDs     []uint `json:"genre_ids" validate:"omitempty,unique,dive,gt=0"`
	MechanicIDs  []uint `json:"mechanic_ids" validate:"omitempty,unique,dive,gt=0"`
	BitIDs       []uint `json:"bit_ids" validate:"omitempty,unique,dive,gt=0"`
	BasegameID   *uint  `json:"basegame_id,omitempty" validate:"omitempty,gt=0"`
	ExpansionIDs []uint `json:"expansion_ids" validate:"omitempty,unique,dive,gt=0"`
}

type UpdateBoardgameRequest struct {
	Title       *string `json:"title" validate:"omitempty"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url" validate:"url"`
}

type BoardgameResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	//Genre        internal.GameGenre `json:"genre"`
	Complexity   int                      `json:"complexity"`
	Learnability int                      `json:"learnability"`
	MinPlayers   int                      `json:"players_min"`
	MaxPlayers   int                      `json:"players_max"`
	BestPlayers  int                      `json:"players_best"`
	Playtime     int                      `json:"playtime"`
	Designer     string                   `json:"designer"`
	PublisherID  uint                     `json:"publisher_id"`
	ImageURL     string                   `json:"image_url"`
	ParentID     int                      `json:"parentid"`
	Genres       []TagResponse            `json:"genres,omitempty"`
	Mechanics    []TagResponse            `json:"mechanics,omitempty"`
	Expansions   []BoardgameTerseResponse `json:"expansions,omitempty"`
}

type BoardgameTerseResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"string"`
}
