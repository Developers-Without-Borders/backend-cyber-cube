package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

func updateScore(user *User, newScore float64) (User, error) {
	client, ctx, cancel := getConnection()
	dbName := os.Getenv("MONGODB_DBNAME")
	defer cancel()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, ctx)

	info, err := GetUserInfo(user)
	if err != nil {
		return *user, err
	}
	filter := bson.M{"_id": bson.M{"$eq": info.ID}}
	update := bson.M{
		"$set": bson.M{
			"username":   info.Username,
			"device_id":  info.DeviceId,
			"score":      newScore,
			"updated_at": time.Now(),
		},
	}
	info.Score = newScore
	_, err = client.Database(dbName).Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return info, err
	}

	return info, nil
}
