package claudinary

import (
	"context"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/echo/v4"
)

type Config struct {
	CLOUDNAME  string
	APIKEY     string
	SECRETKEY  string
	FOLDERNAME string
}

var conf Config

func GetFile(ctx echo.Context) interface{} {
	//upload
	formHeader, err := ctx.FormFile("file")
	if err != nil {
		return err
	}
	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		return err
	}
	return formFile

}
func ImageUploadHelper(input interface{}, folderDetail string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	conf = Config{
		CLOUDNAME:  os.Getenv("CLOUD_NAME"),
		APIKEY:     os.Getenv("API_KEY"),
		SECRETKEY:  os.Getenv("API_SECRET"),
		FOLDERNAME: os.Getenv("UPLOAD_FOLDER"),
	}

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(conf.CLOUDNAME, conf.APIKEY, conf.SECRETKEY)
	if err != nil {
		return "", err
	}
	folder := conf.FOLDERNAME + "/" + folderDetail

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: folder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
