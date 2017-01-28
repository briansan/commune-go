package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/briansan/commune-go/api"
	"github.com/briansan/commune-go/www"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	api.Init(e)
	www.Init(e)
	e.Start(":8888")
}
