package types

import "time"

type StudentReq struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	ClassName string `json:"class_name,omitempty"`
}

type StudentSearchReq struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ClassName string `json:"class_name"`
}

type StudentAddReq struct {
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name" validate:"required,min=2"`
	ClassName string `json:"class_name"`
}

type Student struct {
	ID        int        `json:"id" bson:"id"`
	FirstName string     `json:"first_name" bson:"first_name"`
	LastName  string     `json:"last_name" bson:"last_name"`
	Email     string     `json:"email" bson:"email"`
	ClassName string     `json:"class_name" bson:"class_name"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
}

type DeleteReq struct {
	ID int `json:"id"`
}
