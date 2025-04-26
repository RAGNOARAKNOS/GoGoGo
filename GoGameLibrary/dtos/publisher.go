package dtos

type NewPublisherRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePublisherRequest struct {
	Name string `json:"name,omitempty"`
}

type PublisherResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	//Games []BoardgameResponse `json:"games,omitempty"`
}
