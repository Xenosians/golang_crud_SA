package main

import (
	"github.com/Xenosians/golang_crud_SA/controllers"
	"github.com/Xenosians/golang_crud_SA/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVAriables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/get", controllers.PostsIndex)
	r.GET("/getById/:id", controllers.PostsShow)
	r.PUT("/update/:id", controllers.PostUpdate)
	r.DELETE("/delete/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
