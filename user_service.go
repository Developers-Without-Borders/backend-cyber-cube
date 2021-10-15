package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"time"
)

func Create(user *User) (primitive.ObjectID, error) {
	dbName := os.Getenv("MONGODB_DBNAME")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, ctx)
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	fmt.Println("*********************USER*******************************")
	fmt.Println(user.Username)
	fmt.Println(user.DeviceId)
	fmt.Println("****************************************************")

	result, err := client.Database(dbName).Collection("users").InsertOne(ctx, user)
	if err != nil {
		log.Printf("Could not create Task: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func GetUserInfo(user *User) (User, error) {
	result := User{}
	client, ctx, cancel := getConnection()
	dbName := os.Getenv("MONGODB_DBNAME")
	filter := bson.D{primitive.E{Key: "username", Value: user.Username}}
	defer cancel()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, ctx)
	err := client.Database(dbName).Collection("users").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return User{}, nil
	}

	if result.DeviceId != user.DeviceId {
		return *user, errors.New("username mismatched with this device")
	}
	if err != nil {
		return result, errors.New("username doesnt exists")
	}
	return result, nil
}

func GetUserInfoByDeviceId(user *User) (UserResponse, error) {
	result := UserResponse{}
	client, ctx, cancel := getConnection()
	dbName := os.Getenv("MONGODB_DBNAME")
	filter := bson.D{primitive.E{Key: "device_id", Value: user.DeviceId}}
	defer cancel()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
		}
	}(client, ctx)
	empty := UserResponse{}
	err := client.Database(dbName).Collection("users").FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		empty.IsSuccess = false
		return empty, nil
	}
	empty.IsSuccess = true
	return result, nil
}

func deleteInfoByDeviceId(deviceId string) (int64, error) {

	client, ctx, cancel := getConnection()
	dbName := os.Getenv("MONGODB_DBNAME")
	filter := bson.D{primitive.E{Key: "device_id", Value: deviceId}}
	defer cancel()
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, ctx)
	result, err := client.Database(dbName).Collection("users").DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return 0, err

	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return result.DeletedCount, nil
}
