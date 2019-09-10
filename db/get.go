package db

import (
	"context"
	"time"

	"github.com/letanthang/my_framework/db/types"
)

func GetOneStudent() (*types.Student, error) {
	var student types.Student
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := client.Database(dbName).Collection("student").FindOne(ctx, struct{}{}).Decode(&student)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func GetStudent() (*[]types.Student, error) {
	var students []types.Student
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := client.Database(dbName).Collection("student").Find(ctx, struct{}{})

	if err != nil {
		return nil, err
	}

	for cur.Next(nil) {
		var student types.Student
		err = cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}
