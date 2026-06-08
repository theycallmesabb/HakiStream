package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"hakistream.com/config"
	"hakistream.com/models"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {

	var reg RegisterRequest
	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return

	}
	hasedpass, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{
		Name:      reg.Name,
		Email:     reg.Email,
		Password:  string(hasedpass),
		CreatedAt: time.Now(),
	}
	_, err = config.Db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	//check for if the email already exists
	var userexist models.User
	err = config.Db.Collection("users").FindOne(context.TODO(), bson.M{"email": reg.Email}).Decode(&userexist)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "email already registered",
		})
	}
	//Everthing is okay
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered",
	})

}
