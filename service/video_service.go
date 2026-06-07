package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeFile(c *gin.Context) {
	id := c.Param("id")
	basepath := "./uploads"
	path := filepath.Join(basepath, id)

	if _, err := os.Stat(path); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "The file you are looking for is not found",
		})
		return
	}

	file, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	filesize := fileinfo.Size()

	rangeheader := c.GetHeader("Range")
	if rangeheader == "" {
		c.File(path)
		return
	}
	trimmed := strings.TrimPrefix(rangeheader, "bytes=")
	parts := strings.Split(trimmed, "-")
	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var end int64
	if parts[1] == "" {
		end = filesize - 1
	} else {
		end, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	_, err = file.Seek(start, io.SeekStart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	contentrange := fmt.Sprintf("bytes %d-%d/%d", start, end, filesize)
	c.Header("Content-Range", contentrange)
	c.Header("Content-Length", fmt.Sprintf("%d", end+1-start))
	c.Header("Content-Type", "video/mp4")
	c.Header("Accept-Ranges", "bytes")
	c.Status(http.StatusPartialContent)
	log.Println()
	_, err = io.CopyN(c.Writer, file, end+1-start)
	if err != nil && err != io.EOF {
		return
	}
}
