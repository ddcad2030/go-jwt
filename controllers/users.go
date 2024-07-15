package controllers

import (
	"david/go-jwt/initial"
	"david/go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to hash password",
		})
		return
	}

	post := models.User{Email: body.Email, Password: string(hash)}
	result := initial.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to create",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create user successful",
	})
}
