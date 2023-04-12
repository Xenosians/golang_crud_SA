package main

import (
	"github.com/Xenosians/golang_crud_SA/initializers"
	"github.com/Xenosians/golang_crud_SA/models"
)

func init() {
	initializers.LoadEnvVAriables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
