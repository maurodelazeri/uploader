package api

import (
	"errors"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

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

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func Upload(c echo.Context) error {
	req := c.Request()
	headers := req.Header
	bucketName := headers.Get("bucketName")
	objectName := headers.Get("objectName")

	if len(bucketName) == 0 || len(objectName) == 0 {
		return errors.New("bucketName or objectName empty")
	}

	s := strings.Split("/storage/"+bucketName+"/"+objectName, "/")
	path := ""
	index := 0
	for index < len(s)-1 {
		path += s[index] + "/"
		index++
	}

	err := makeDirectoryIfNotExists(path)
	if err != nil {
		return errors.New("err creating DIR " + err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		return errors.New("err 0" + err.Error())
	}
	data, err := readFileBody(file)
	if err != nil {
		return errors.New("err 1" + err.Error())
	}

	f, err := os.Create("/storage/" + bucketName + "/" + objectName)
	if err != nil {
		return errors.New("err 2 " + "/storage/" + bucketName + "/" + objectName + " " + err.Error())
		//return err
	}
	defer f.Close()
	f.Write(data)
	return c.JSON(http.StatusOK, "ok")
}
