package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/service"
)

func FoodRouter(e *echo.Echo) {
	// เรียกใช้งานฟังก์ชัน service
	e.GET("/menu", service.GetMenu)
	e.GET("/menu/:foodId", service.GetMenuById)
	e.GET("/menu/categorys", service.GetCategory)
	e.GET("/menu/category/:categoryId", service.GetMenuByCategory)
}
