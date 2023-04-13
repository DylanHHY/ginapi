package main

import (
	"goAPI/initializers"
	model "goAPI/models"
)

func init() {
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}
