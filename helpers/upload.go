package helpers

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	uuid "github.com/google/uuid"
)

func CloudUpload(src interface{}) (string, error) {
	cloud := os.Getenv("CLOUD_NAME")
	key := os.Getenv("CLOUD_KEY")
	secret := os.Getenv("CLOUD_SECRET")

	ctx := context.Background()
	imageID := uuid.New()

	cld, err := cloudinary.NewFromParams(cloud, key, secret)
	if err != nil {
		return "", err
	}

	res, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		PublicID: imageID.String(),
		Folder:   "vehicle",
	})
	if err != nil {
		return "", err
	}

	image, err := cld.Image("sample.jpg")
	if err != nil {
		return "", err
	}

	image.Transformation = "c_fill,h_150,w_100"

	return res.SecureURL, nil
}
