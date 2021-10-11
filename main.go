package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
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
