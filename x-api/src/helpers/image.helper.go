package helpers

import (
	"io"
	"os"

	"github.com/apot-group/golang-skeleton/x-api/src/settings"
)

func UploadImage(fileName string, file io.Reader) (path string, err error) {
	filePath := settings.ConfEnv.Storage.Image + fileName
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return filePath[1:], nil

}
