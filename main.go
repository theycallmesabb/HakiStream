package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hakistream.com/config"
	"hakistream.com/routes"
)

func main() {
	fmt.Println("Server started: ")

	r := gin.Default()
	godotenv.Load()
	routes.SetupRoutes(r)
	config.ConnectDb()
	config.ConnectRedis()
	config.ConnectR2()
	r.Run(":8080")
}
