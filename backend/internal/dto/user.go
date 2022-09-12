package dto

type User struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required"`
	ID        string `json:"id" validate:"required"`
}
