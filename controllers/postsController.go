package controllers

import (
	"github.com/Xenosians/golang_crud_SA/initializers"
	"github.com/Xenosians/golang_crud_SA/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//req body
	var body struct {
		Body       string
		StoryTitle string
	}
	c.Bind(&body)
	//create post
	post := models.Post{StoryTitle: body.StoryTitle, Body: body.Body}
	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}
	//return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get Post
	var posts []models.Post
	initializers.DB.Find(&posts)
	//Respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Body       string
		StoryTitle string
	}
	c.Bind(&body)
	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{
		StoryTitle: body.StoryTitle,
		Body:       body.Body,
	})
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
