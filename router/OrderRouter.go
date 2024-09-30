package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/service"
)

func OrderRouter(e *echo.Echo) {
	e.GET("/orders", service.GetOrders)
	e.PUT("/order/cancel/:order_id", service.CancelOrder)
	e.PUT("/order/paid/:order_id", service.ChackBillOrder)
}
