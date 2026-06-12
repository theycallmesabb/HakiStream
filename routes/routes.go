package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hakistream.com/handlers"
	"hakistream.com/middleware"
	"hakistream.com/service"
)

func SetupRoutes(r *gin.Engine) string {
	// only 50 mb allowed to upload
	r.MaxMultipartMemory = 500 << 20

	tokenzied := r.Group("/")
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.GET("/ui", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.GET("/gallery", func(c *gin.Context) {
		c.File("./static/gallery.html")
	})

	//this are the routes which are protected
	tokenzied.Use(middleware.AuthMiddleware())
	tokenzied.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is up",
		})
	})

	tokenzied.POST("/uploadmov", service.UploadMovie)

	tokenzied.GET("/videos/:id", service.ServeFile)
	tokenzied.GET("/videos", service.ListVideos)

	tokenzied.POST("/logout", handlers.LogOut)

	return "ok"
}
