package dtos

type NewTagRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateTagRequest struct {
	Name string `json:"name,omitempty"`
}

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
