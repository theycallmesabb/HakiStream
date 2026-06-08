package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hakistream.com/handlers"
	"hakistream.com/service"
)

func SetupRoutes(r *gin.Engine) string {
	// only 50 mb allowed to upload
	r.MaxMultipartMemory = 500 << 20
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is up",
		})
	})

	r.POST("/uploadmov", service.UploadMovie)

	r.GET("/videos/:id", service.ServeFile)
	r.GET("/videos", service.ListVideos)
	r.GET("/ui", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.POST("/register", handlers.RegisterUser)
	return "ok"

}
