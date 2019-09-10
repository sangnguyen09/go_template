package db

import (
	"context"
	"time"

	"github.com/letanthang/my_framework/db/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOneStudent() (*types.Student, error) {
	var student types.Student
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := Client.Database(dbName).Collection("student").FindOne(ctx, struct{}{}).Decode(&student)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func GetStudent() (*[]types.Student, error) {
	var students []types.Student
	findOptions := options.Find()
	findOptions.SetLimit(30)
	cur, err := Client.Database(dbName).Collection("student").Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var student types.Student
		err = cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}
