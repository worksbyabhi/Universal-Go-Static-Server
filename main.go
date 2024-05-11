package main

import (
	"embed"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed web
var webStatic embed.FS

func main() {
	e := echo.New()

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "web",
		HTML5:      true,
		Filesystem: http.FS(webStatic),
	}))

	e.GET("/", func(c echo.Context) error {
		return c.File("web/index.html")
	})

	err := e.Start(":8001")

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panic(err)
	}
}
