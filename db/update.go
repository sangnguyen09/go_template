package db

import (
	"context"
	"time"

	"github.com/letanthang/my_framework/db/types"
)

func InsertStudent(req types.StudentAddReq) (interface{}, error) {
	student := types.Student{FirstName: req.FirstName, LastName: req.LastName, ClassName: req.ClassName}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := db.Collection("student").InsertOne(ctx, student)
	id := res.InsertedID
	return id, err
}
