package helpers

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(file multipart.File, fileName string, destination string) error {
	dst, err := os.Create(destination + "/" + fileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the file data to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}
