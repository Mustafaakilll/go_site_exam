package models

type PaginateRequest struct {
	Limit  int `json:"limit" validate:"max=100,min=0"`
	Offset int `json:"offset"`
}
