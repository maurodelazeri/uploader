package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/quiknode-labs/uploader/api"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/upload", api.Upload)
	e.Logger.Fatal(e.Start(":1323"))
}
