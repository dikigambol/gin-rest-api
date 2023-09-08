package controllers

import (
	"gin-rest-api/initializers"
	"gin-rest-api/models"

	"github.com/gin-gonic/gin"
)

// create
func PostCreate(c *gin.Context) {
	// get req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
	})
}

// read
func PostIndex(c *gin.Context) {
	// get posts data
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// read 1 data
func PostShow(c *gin.Context) {
	// get from url
	id := c.Param("id")

	// get post data
	var post []models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

// update 
func PostUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data from body req
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// show data
	var post []models.Post
	initializers.DB.First(&post, id)

	// update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respons
	c.JSON(200, gin.H{
		"status": 200,
	})
}

// delete 
// read 1 data
func PostDelete(c *gin.Context) {
	// get from url
	id := c.Param("id")

	// delete
	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"status": 200,
	})
}
