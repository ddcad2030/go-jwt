package main

import (
	"david/go-jwt/controllers"
	"david/go-jwt/initial"
	"david/go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initial.LoadEnv()
	initial.ConnectDB()
	initial.SyncDatabase()
}
func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.Auth, controllers.Validate)
	r.Run()
}
