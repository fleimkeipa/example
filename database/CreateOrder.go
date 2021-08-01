package database

import (
	"go.mongodb.org/mongo-driver/bson"
)

func Create(v interface{}, collectionName string) error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	data, err := bson.Marshal(v)
	if err != nil {
		return err
	}
	collection := client.Database("tesodev").Collection(collectionName)
	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}