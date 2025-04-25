package dtos

type NewMechanicRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateMechanicRequest struct {
	Name string `json:"name,omitempty"`
}

type MechanicResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
