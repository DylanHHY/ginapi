package controllers

import (
	"goAPI/initializers"
	model "goAPI/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// create a post
	post := model.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []model.Post
	initializers.DB.Find(&posts)
	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Get the posts
	var post model.Post
	initializers.DB.First(&post, id)
	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the data odd req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post model.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(model.Post{
		Title: body.Title,
		Body:  body.Body})

	// Respond with it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&model.Post{}, id)
	// Respond
	c.Status(200)
}
