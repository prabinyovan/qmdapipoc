package app

import (
	"github.com/labstack/echo/v4"
)

var (
	router = echo.New()
)

func StartApplication() {

	mapUrls()
	router.Start(":8081")
}
