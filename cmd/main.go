// cmd/main.go
package main

import (
	"github.com/labstack/echo/v4"
	"restaurant/router"
)

func main() {
	e := echo.New()

	router.CustomerRouter(e)
	router.FoodRouter(e)
	router.OrderRouter(e)
	e.Logger.Fatal(e.Start(":8080"))
}
