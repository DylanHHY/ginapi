package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/rooms", func(c *gin.Context) {
		res := []string{"foo", "bar"}
		c.JSON(200, res)
	})
	router.POST("/rooms", func(c *gin.Context) {
		// id := c.Param("id")
		res := []string{"Hello!!"}
		c.JSON(200, res)
	})
	router.POST("/rooms/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	router.PUT("/rooms/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	router.DELETE("/rooms/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
