package main

import (
	"github.com/ceramicyu/prepare/echo_usage/controller"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo)*echo.Echo{
	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		writeCookie(cc)
		return cc.String(200, "O22K")
	})
	e.GET("/user", controller.User)
	return e
}
