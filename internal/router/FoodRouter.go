package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/internal/service"
)

func FoodRouter(e *echo.Echo) {
	e.GET("food", service.GetMenu)
	e.GET("food/:foodId", service.GetMenuById)
	e.GET("food/categorys", service.GetCategory)
	e.GET("food/category/:categoryId", service.GetMenuByCategory)
	e.POST("food/create", service.CreateFood)
	e.POST("food/create/category", service.CreateCategory)
	e.PUT("food/edit/:foodId", service.EditFood)
	e.PUT("food/category/edit/:categoryId", service.CategoryEdit)
	e.PUT("food/available/:status/:foodId", service.AvailableFood)
	e.DELETE("food/delete/:foodId", service.DeleteFood)
}
