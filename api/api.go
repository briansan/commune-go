package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) error {

	// setup /api
	api := e.Group("/api")

	// setup /api/service
	svc := api.Group("/service")

	// ping pong
	svc.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// setup the rest
	return initV1(api)
}
