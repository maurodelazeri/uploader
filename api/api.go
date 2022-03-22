package api

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
)

func readFileBody(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		log.Println("Error2:", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(src)
	if err != nil {
		log.Println("Error3:", err)
		return nil, err
	}
	return body, nil
}

func Upload(c echo.Context) error {
	req := c.Request()
	headers := req.Header
	bucketName := headers.Get("bucketName")
	objectName := headers.Get("objectName")
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("problem to read from file:", err)
		return err
	}
	data, err := readFileBody(file)
	if err != nil {
		return err
	}
	f, err := os.Create("/storage/" + bucketName + "/" + objectName)
	if err != nil {
		log.Println("problem create storage file:", err)
		return err
	}
	defer f.Close()
	f.Write(data)
	return c.JSON(http.StatusOK, `{"status":"ok"}`)
}
