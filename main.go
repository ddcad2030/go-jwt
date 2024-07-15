package main

import (
	"david/go-jwt/controllers"
	"david/go-jwt/initial"

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
	r.Run()
}
