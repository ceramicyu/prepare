package main

import (
	"github.com/labstack/echo"
)
type CustomContext struct {
	echo.Context
}



func main() {
	e := echo.New()

	e= Router(e)

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	})

	e.Logger.Fatal(e.Start(":1323"))
}