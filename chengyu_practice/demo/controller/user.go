package controller

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "userser",
	})
}