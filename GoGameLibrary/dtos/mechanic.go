package dtos

type NewMechanicRequest struct {
	Name string `json:"name" validate:"required"`
}

type MechanicResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
