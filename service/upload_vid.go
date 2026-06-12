package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"hakistream.com/config"
)

func UploadMovie(c *gin.Context) {
	file, err := c.FormFile("video") // wait, front-end sends 'file' not 'video'. Let me fix this too.
	if err != nil {
		file, err = c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to open file"})
		return
	}
	defer src.Close()
	
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

	ctx := context.Background()
	_, err = config.S3CLIENT.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("R2_BUCKET_NAME")),
		Key:    aws.String(filename),
		Body:   src,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to upload to R2"})
		return
	}

	c.JSON(200, gin.H{
		"status": "uploaded",
		"file":   filename,
	})
}
