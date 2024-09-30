package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/service"
)

func EmployeeRouter(e *echo.Echo) {
	e.GET("/employees", service.GetEmployees)
}
