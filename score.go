package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
const NegativeValueError="NEGATIVE_VALUE_ERROR"
func handleUpdateScore(c *gin.Context){
	var user User
	user.Username = c.Param("name")
	user.DeviceId = c.Param("deviceId")
	newScore := c.Param("score")
	if s, err := strconv.ParseFloat(newScore, 64); err == nil {
		if s < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "negative value error","code":NegativeValueError })
			return
		}
		info, err := updateScore(&user,s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error() })
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": info})
	}
}