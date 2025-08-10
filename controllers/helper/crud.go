package helper

import (
	"context"
	"intern_backend/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RetrieveData(filter bson.M, doc string, obj interface{}) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetDatabase().Collection(doc)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil

}

func RetrieveOneData(filter bson.M, doc string, obj interface{}) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetDatabase().Collection(doc)
	err := collection.FindOne(ctx, filter).Decode(obj)
    if err != nil {
        return nil, err
    }

    return obj, nil

}

func InsertData(doc string, obj interface{}) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// err := utils.GetValidator().Struct(obj)
	// if err != nil {
	// 	return nil, err
	// }

	collection := config.GetDatabase().Collection(doc)
	res, err := collection.InsertOne(ctx, obj)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func UpdateData(doc string, key string, value interface{}, obj interface{}) (*mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetDatabase().Collection(doc)
	filter := bson.M{key: value}
	update := bson.M{"$set": obj}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func DeleteData(doc string, key string, value interface{}) (*mongo.DeleteResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetDatabase().Collection(doc)
	filter := bson.M{key: value}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return res, nil

}
