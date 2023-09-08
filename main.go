package main

import (
	"gin-rest-api/controllers"
	"gin-rest-api/initializers"
	"gin-rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.POST("/posts", middleware.AuthValidate, controllers.PostCreate)
	r.GET("/posts", middleware.AuthValidate, controllers.PostIndex)
	r.GET("/posts/:id", middleware.AuthValidate, controllers.PostShow)
	r.PUT("/posts/:id", middleware.AuthValidate, controllers.PostUpdate)
	r.DELETE("/posts/:id", middleware.AuthValidate, controllers.PostDelete)

	r.Run()
}
