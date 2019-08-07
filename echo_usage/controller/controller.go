package controller

import (
	"encoding/json"
	"github.com/labstack/echo"
	"time"
)



func  User(e echo.Context) error {


		type User struct {
			Name string `json:"name"`
			Age int `json:"age"`
		}
		time.Sleep(84*time.Second)
		bytes,_:=json.Marshal(User{Name:"john",Age:455})
		//e.String(200, string(bytes))

	return e.String(200, string(bytes))
}