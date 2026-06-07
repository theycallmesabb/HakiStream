package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"hakistream.com/routes"
)

func main() {
	fmt.Println("Server started: ")

	r := gin.Default()

	routes.SetupRoutes(r)
	r.Run(":8080")
}
