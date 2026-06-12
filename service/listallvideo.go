package service

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"hakistream.com/config"
)

type Videos struct {
	Video string `json:"name"`
	URL   string `json:"url"`
}

func ListVideos(c *gin.Context) {
	bucket := os.Getenv("R2_BUCKET_NAME")
	output, err := config.S3CLIENT.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to list videos from r2: " + err.Error(),
		})
		return
	}

	vid := []Videos{}
	for _, object := range output.Contents {
		name := *object.Key
		vid = append(vid, Videos{
			Video: name,
			URL:   "/videos/" + name,
		})
	}

	c.JSON(http.StatusOK, vid)
}
