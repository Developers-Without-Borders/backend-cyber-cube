package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Username  string             `bson:"username"`
	DeviceId  string             `bson:"device_id"`
	Score     float64            `bson:"score"`
}

const UsernameExist = "USERNAME_EXIST"
const SUCCESS = "SUCCESS"

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ping alive"})
}

func handleGetUserInfo(c *gin.Context) {
	var user User
	user.Username = c.Param("name")
	user.DeviceId = c.Param("deviceId")
	info, err := GetUserInfo(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": info})
}

func handleCreateUser(c *gin.Context) {
	var user User
	user.Username = c.Param("name")
	user.DeviceId = c.Param("deviceId")
	result, err := GetUserInfo(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if result.Username != "" && result.DeviceId != "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "username already exists", "code": UsernameExist})
		return
	}

	user.Score = 0.0
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	id, err := Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func handleGetUserInfoByDeviceId(c *gin.Context) {
	var user User
	user.DeviceId = c.Param("deviceId")
	info, err := GetUserInfoByDeviceId(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": info})
}
