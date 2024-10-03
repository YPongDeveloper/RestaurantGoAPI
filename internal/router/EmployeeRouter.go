package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/internal/service"
)

func EmployeeRouter(e *echo.Echo) {
	e.GET("employees", service.GetEmployees)
	e.GET("employee/:employeeId", service.GetEmployeeById)
	e.POST("employee/hire", service.CreateEmployee)
	e.PUT("employee/edit/:employeeId", service.UpdateEmployee)
	e.PUT("employee/fire/:employeeId", service.FireEmployee)
	e.PUT("employee/absent/:employeeId", service.AbsentEmployee)
	e.PUT("employee/work/:employeeId", service.WorkEmployee)

}
