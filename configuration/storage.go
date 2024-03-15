package configuration

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	Bucket *string
	Uploader *manager.Uploader
}

func NewStorage(config Config) *Storage {
	client := s3.New(s3.Options{
		Region: "ap-southeast-1",
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.Get("S3_ID"), config.Get("S3_SECRET_KEY"), "")),
	})

	uploader := manager.NewUploader(client)

	return &Storage{
		Bucket: aws.String(config.Get("S3_BUCKET_NAME")),
		Uploader: uploader,
	}
}