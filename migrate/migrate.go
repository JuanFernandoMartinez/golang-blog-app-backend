package main

import (
	"example.com/blog-app/initializers"
	"example.com/blog-app/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Thought{})
	initializers.DB.AutoMigrate(&models.Comment{})
}
