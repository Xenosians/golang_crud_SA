package main

import (
	"github.com/Xenosians/golang_crud_SA/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVAriables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
