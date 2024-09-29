package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"restaurant/model/request"
	"restaurant/model/response"
	"restaurant/repository"
	"strconv"
)

func CancelOrder(c echo.Context) error {
	// รับค่า order_id จาก path parameter
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid order ID"))
	}
	var reviewRequest request.ReviewRequest
	if err := c.Bind(&reviewRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request body"))
	}
	// เรียกใช้ฟังก์ชัน CancelOrderData
	err = repository.UpdateStatusOrderData(orderId, 3, reviewRequest.Review)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, response.ErrorResponse("Order not found"))
		}
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to cancel order"))
	}

	// ส่ง response สำเร็จกลับไป
	return c.JSON(http.StatusOK, response.SuccessResponse("Order canceled successfully"))
}

func ChackBillOrder(c echo.Context) error {
	// รับค่า order_id จาก path parameter
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid order ID"))
	}
	var reviewRequest request.ReviewRequest
	if err := c.Bind(&reviewRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request body"))
	}
	// เรียกใช้ฟังก์ชัน UpdateStatusOrderData
	err = repository.UpdateStatusOrderData(orderId, 4, reviewRequest.Review)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, response.ErrorResponse("Order not found"))
		}
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to check bill order"))
	}

	// ส่ง response สำเร็จกลับไป
	return c.JSON(http.StatusOK, response.SuccessResponse("Order checked bill successfully"))
}
