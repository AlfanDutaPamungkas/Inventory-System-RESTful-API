package helper

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(ctx context.Context, cld *cloudinary.Cloudinary, file multipart.File, fileHeader *multipart.FileHeader) string {
	imgParam := uploader.UploadParams{
		PublicID:       fileHeader.Filename,
		Folder:         "inventory",
		AllowedFormats: []string{"jpg", "png", "jpeg"},
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, imgParam)
	PanicError(err)

	return uploadResult.SecureURL
}
