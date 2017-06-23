package errors

import (
	"github.com/labstack/echo"
)

type HTTPError interface {
	error
	StatusCode() int
}

func Response(c echo.Context, err HTTPError) error {
	return c.JSON(err.StatusCode(), map[string]string{"err": err.Error()})
}
