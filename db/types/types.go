package types

import "time"

type StudentAddReq struct {
	FirstName string `json:"first_name" validate:"required,min=3"`
	LastName  string `json:"last_name" validate:"required"`
	ClassName string `json:"class_name"`
}

type Student struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string
	Email     string
	ClassName string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type DeleteReq struct {
	ID int `json:"id"`
}
