package app

import (
	"inventory-system-api/helper"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func NewCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	helper.PanicError(err)
	return cld
}
