package main

import (
	"github.com/gin-gonic/gin"

)


func main() {
	r := gin.Default()
	r.GET("/user/:name/:deviceId" , handleGetUserInfo)
	r.PUT("/user/:name/:deviceId" ,handleCreateUser)
	r.PUT("/score/:name/:deviceId/:score",handleUpdateScore)

	//client := redis.NewClient(&redis.Options{
	//	Addr: "localhost:6379",
	//	Password: "",
	//	DB: 0,
	//})
	//
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)

	err := r.Run()
	if err != nil {
		return 
	}
}