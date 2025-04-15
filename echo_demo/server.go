package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main(){
    e := echo.New()

    e.GET("/", func (c echo.Context) error {
        return c.String(http.StatusOK, "Hello world and its awesome")
    })

    e.POST("/users", func (c echo.Context) error{
        return c.String(http.StatusOK, "Users created")
    })


    e.Logger.Fatal(e.Start(":1323"))
}
