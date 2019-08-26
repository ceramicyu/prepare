package main

import (
	"github.com/ceramicyu/prepare/chengyu_practice/demo/controller"
	"github.com/gin-gonic/gin"
)

/*
docker build -t server .

docker run --rm -it -d -p 3000:3000  -v /data:/data  server
 */

func main() {


	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user",controller.GetUserInfo)
	r.Run() // listen and serve on 0.0.0.0:8080

}

