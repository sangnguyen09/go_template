package helpers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Counter struct {
	ID    string `bson:"id"`
	Value int    `bson:"seq_val"`
}

func GetNextID(col *mongo.Collection, sequenceName string) (int, error) {
	filter := bson.M{"id": sequenceName}
	update := bson.M{"$inc": bson.M{"seq_val": 1}}
	option := options.FindOneAndUpdate()
	option.SetUpsert(true)
	var counter Counter
	err := col.FindOneAndUpdate(context.TODO(), filter, update, option).Decode(&counter)
	if err != nil {
		return 1, err
	}
	return counter.Value + 1, err
}
