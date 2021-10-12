package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")
	fmt.Print(username)
	fmt.Print(password)
	fmt.Print(clusterEndpoint)

	r := gin.Default()
	r.GET("/ping", handlePing)
	r.GET("/user/:name/:deviceId", handleGetUserInfo)
	r.GET("/user/info/:deviceId", handleGetUserInfoByDeviceId)
	r.PUT("/user/:name/:deviceId", handleCreateUser)
	r.PUT("/score/:name/:deviceId/:score", handleUpdateScore)
	r.DELETE("/user/:deviceId", handleDeleteUserInfo)

	err := r.Run()
	if err != nil {
		return
	}
}
