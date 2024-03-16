package service

import (
	"context"
	"mime/multipart"
)

type FileService interface {
	Upload(ctx context.Context, file *multipart.FileHeader) string
}
