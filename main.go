package main

import (
	"goAPI/controllers"
	"goAPI/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("posts", controllers.PostsIndex)
	r.GET("posts/:id", controllers.PostsShow)
	r.PUT("posts/:id", controllers.PostsUpdate)
	r.DELETE("posts/:id", controllers.PostsDelete)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
