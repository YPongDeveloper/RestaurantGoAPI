package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/service"
)

func FoodRouter(e *echo.Echo) {
	e.GET("menu", service.GetMenu)
	e.GET("menu/:foodId", service.GetMenuById)
	e.GET("menu/categorys", service.GetCategory)
	e.GET("menu/category/:categoryId", service.GetMenuByCategory)
	e.POST("food/create", service.CreateFood)
	e.POST("food/create/category", service.CreateCategory)
	e.PUT("food/edit/:foodId", service.EditFood)
	e.PUT("food/available/:status/:foodId", service.AvailableFood)
	e.DELETE("food/delete/:foodId", service.DeleteFood)
}
