package router

import (
	"github.com/labstack/echo/v4"
	"restaurant/service"
)

func OrderRouter(e *echo.Echo) {
	e.GET("orders", service.GetOrders)
	e.GET("order/:orderId", service.GetOrderById)
	e.GET("order/wait", service.GetOrderWait)
	e.GET("order/eating", service.GetOrderEating)
	e.GET("order/paid", service.GetOrderPaid)
	e.GET("order/cancel", service.GetOrderCancel)
	e.GET("order/dateHistory/:date", service.GetOrderDateHistory)
	e.GET("order/employee/:employeeId", service.GetEmployeeIdOrder)
	e.PUT("order/cancel/:order_id", service.CancelOrder)
	e.PUT("order/eating/:order_id", service.EatingOrder)
	e.PUT("order/paid/:order_id", service.ChackBillOrder)
	e.POST("order/create", service.CreateOrder)
}
