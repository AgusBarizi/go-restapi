package dto

type CategoryCreateRequest struct {
	Name string `validate:"required,max=75,min=1" json:"name"`
}
