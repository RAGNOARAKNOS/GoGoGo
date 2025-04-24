package dtos

type NewPublisherRequest struct {
	Name string `json:"name" validate:"required"`
}

type PublisherResponse struct {
	ID    uint                `json:"id"`
	Name  string              `json:"name"`
	Games []BoardgameResponse `json:"games,omitempty"`
}
