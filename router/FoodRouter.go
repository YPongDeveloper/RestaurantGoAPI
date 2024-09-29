package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/service"
)

func FoodRouter(e *echo.Echo) {
	// เรียกใช้งานฟังก์ชัน service
	e.GET("/menu", service.GetMenu)
}
