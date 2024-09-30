// cmd/router.go
package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/service"
)

func CustomerRouter(e *echo.Echo) {
	// เรียกใช้งานฟังก์ชัน service
	e.GET("/customers", service.GetCustomers)
	e.GET("/customer/:customerId", service.GetCustomerById)
}
