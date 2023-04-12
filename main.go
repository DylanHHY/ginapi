package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "tom")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			// "success": "true",
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "tom")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	r.PUT("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "tom")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	r.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
