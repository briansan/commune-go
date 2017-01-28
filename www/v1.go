package www

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	// setup /www
	g := e.Group("/www")
	g.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	g.File("/upload", "www/public/index.html")
}
