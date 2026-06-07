package service

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadMovie(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ext := filepath.Ext(file.Filename)
	filename := id + ext

	path := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(500, gin.H{"error": "failed to save file"})
		return
	}

	c.JSON(200, gin.H{
		"status": "uploaded",
		"file":   filename,
	})
}
