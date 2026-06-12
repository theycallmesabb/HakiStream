package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"hakistream.com/config"
)

func ServeFile(c *gin.Context) {
	id := c.Param("id")
	bucket := os.Getenv("R2_BUCKET_NAME")
	ctx := context.Background()

	// Prepare S3 GetObject Input
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(id),
	}

	// Forward Range header if it exists
	rangeHeader := c.GetHeader("Range")
	if rangeHeader != "" {
		input.Range = aws.String(rangeHeader)
	}

	output, err := config.S3CLIENT.GetObject(ctx, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The file you are looking for is not found or inaccessible in R2: " + err.Error(),
		})
		return
	}
	defer output.Body.Close()

	// Forward response headers from R2 to the client
	if output.ContentRange != nil {
		c.Header("Content-Range", *output.ContentRange)
	}
	if output.ContentLength != nil {
		c.Header("Content-Length", fmt.Sprintf("%d", *output.ContentLength))
	} else if output.ContentLength != nil {
		c.Header("Content-Length", fmt.Sprintf("%d", *output.ContentLength))
	}
	
	contentType := "video/mp4"
	if output.ContentType != nil {
		contentType = *output.ContentType
	}
	c.Header("Content-Type", contentType)
	c.Header("Accept-Ranges", "bytes")

	// Check if this is a partial content response
	status := http.StatusOK
	if rangeHeader != "" {
		status = http.StatusPartialContent
	}
	c.Status(status)

	// Stream the video chunks directly to the client
	_, err = io.Copy(c.Writer, output.Body)
	if err != nil && err != io.EOF {
		log.Println("Error streaming from R2:", err)
		return
	}
}
