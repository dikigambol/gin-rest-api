package main

import (
	"gin-rest-api/controllers"
	"gin-rest-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDb()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.PostCreate)

	r.Run()
}
