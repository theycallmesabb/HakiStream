package handlers

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
type LoginRequest struct {
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

	// 1. Check if user already exists
	var userexist models.User
	err := config.Db.Collection("users").
		FindOne(context.TODO(), bson.M{"email": reg.Email}).
		Decode(&userexist)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "email already registered",
		})
		return
	}

	// 2. Hash password
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	// 3. Create user
	user := models.User{
		Name:      reg.Name,
		Email:     reg.Email,
		Password:  string(hashedpass),
		CreatedAt: time.Now(),
	}

	// 4. Insert user
	_, err = config.Db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered",
	})
}

// thisi s used to for login
func LoginUser(c *gin.Context) {
	var login LoginRequest
	var dbUser models.User
	// make login and password available
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}
	err := config.Db.Collection("users").FindOne(context.TODO(),
		bson.M{"email": login.Email}).Decode(&dbUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//comparing password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"user": "Unauthorized",
		})
		return
	}
	//Making JWT token and sending it to client
	claims := jwt.MapClaims{
		"sub": dbUser.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretkey := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in jwt",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "You are logged in",
		"token":   tokenString,
	})

}

// thisis used for logout
func LogOut(c *gin.Context) {
	ctx := context.Background()
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization empty"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	err := config.Rdb.Set(ctx, tokenString, "blacklisted", 24*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
