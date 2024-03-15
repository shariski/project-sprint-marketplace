package impl

import (
	"context"
	"mime/multipart"
	"project-sprint-marketplace/configuration"
	"project-sprint-marketplace/exception"
	"project-sprint-marketplace/service"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type fileServiceImpl struct {
	storage *configuration.Storage
}

func NewFileServiceImpl(
	storage *configuration.Storage,
) service.FileService {
	return &fileServiceImpl{
		storage: storage,
	}
}

func (fileService *fileServiceImpl) Upload(ctx context.Context, file *multipart.FileHeader) string {
	uploader := fileService.storage.Uploader

	uploadFile, err := file.Open()
	exception.PanicLogging(err)

	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: fileService.storage.Bucket,
		Key:    aws.String(strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename),
		ACL: 		"public-read",
		Body:   uploadFile,
	})

	exception.PanicLogging(err)

	return result.Location
}
