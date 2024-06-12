package main

import (
	"example.com/blog-app/initializers"
	"example.com/blog-app/server"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	server.StartServer()
}
