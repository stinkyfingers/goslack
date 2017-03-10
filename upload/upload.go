package upload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(path, channel, token string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	writer.Close()

	u := fmt.Sprintf("https://slack.com/api/files.upload?token=%s&channels=%s&filename=%s", token, channel, path)
	_, err = http.Post(u, writer.FormDataContentType(), body)
	return err
}
