package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"hakistream.com/config"
	"hakistream.com/routes"
)

func main() {
	fmt.Println("Server started: ")

	r := gin.Default()

	routes.SetupRoutes(r)
	config.ConnectDb()
	r.Run(":8080")
}
