package config

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3CLIENT *s3.Client

func ConnectR2() {
	AccessKey := os.Getenv("R2_ACCESS_KEY_ID")
	SecretKey := os.Getenv("R2_SECRET_ACCESS_KEY")
	Endpoint := os.Getenv("R2_ENDPOINT")
	Creds := credentials.NewStaticCredentialsProvider(AccessKey, SecretKey, "")

	S3CLIENT = s3.NewFromConfig(aws.Config{
		Region:       "auto",
		Credentials:  Creds,
		BaseEndpoint: aws.String(Endpoint),
	})
}
