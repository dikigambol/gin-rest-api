package main

import (
	"gin-rest-api/initializers"
	"gin-rest-api/models"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}