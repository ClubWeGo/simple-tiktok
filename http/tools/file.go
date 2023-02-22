package tools

import (
	"io"
	"mime/multipart"
	"os"
)

// SaveFile saves files to disk
func SaveFile(file *multipart.File, savepath string) error {
	var output *os.File

	output, err := os.OpenFile(savepath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, *file)
	if err != nil {
		return err
	}

	return nil
}
