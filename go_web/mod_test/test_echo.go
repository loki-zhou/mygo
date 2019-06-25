package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"mod_test/utils"
)

func main() {
	fmt.Println("hello Echo")
	utils.PrintText("Hi")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
