package models

// import (
// 	"net/http"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )

// type Videos struct {
// 	Video string `json:"name"`
// }

// func ListVideos(c *gin.Context) {
// 	files, err := os.ReadDir("./uploads")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	var vid []Videos
// 	for _, file := range files {
// 		if !file.IsDir() {
// 			vid = append(vid, Videos{Video: file.Name()})
// 		}
// 	}
// 	c.JSON(http.StatusOK, vid)
// }
