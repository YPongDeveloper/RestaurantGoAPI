package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"restaurant/internal/model/request"
	response2 "restaurant/internal/model/response"
	"restaurant/internal/repository"
	"strconv"
)

func CancelOrder(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid order ID"))
	}
	var reviewRequest request.ReviewRequest
	if err := c.Bind(&reviewRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid request body"))
	}

	err = repository.UpdateStatusOrderData(orderId, 3, reviewRequest.Review)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, response2.ErrorResponse("Order not found Or updated"))
		}
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to cancel order"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse("Order canceled successfully"))
}

func ChackBillOrder(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid order ID"))
	}
	var reviewRequest request.ReviewRequest
	if err := c.Bind(&reviewRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid request body"))
	}

	err = repository.UpdateStatusOrderData(orderId, 4, reviewRequest.Review)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, response2.ErrorResponse("Order not found or updated"))
		}
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to check bill order"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse("Order checked bill successfully"))
}

func GetOrders(c echo.Context) error {

	orderResult := repository.GetOrderData()
	if orderResult == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve All order data"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse(orderResult))
}

func GetOrderById(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid order ID"))
	}

	orderResult := repository.GetOrderByIdData(orderId)
	if orderResult == (response2.OrderIdResponse{}) {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve order id data"))
	}
	orderListResult := repository.GetOrderListOrderData(orderId)
	if orderListResult == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve order list data"))
	}

	var orderResponse response2.OrderResponse
	orderResponse.CustomerId = orderResult.CustomerId
	orderResponse.EmployeeId = orderResult.EmployeeId
	orderResponse.TableId = orderResult.TableId
	orderResponse.OrderId = orderResult.OrderId
	orderResponse.TotalAmount = orderResult.TotalAmount
	orderResponse.OrderDate = orderResult.OrderDate
	orderResponse.Status = orderResult.Status
	orderResponse.Review = orderResult.Review
	for _, orderList := range orderListResult {
		if orderList.OrderId == orderResult.OrderId {
			orderResponse.OrderList = append(orderResponse.OrderList, orderList)
		}
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse(orderResponse))
}

func GetOrderWait(c echo.Context) error {
	orderResult := repository.GetOrderStatustData(1)
	if orderResult == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve order wait data"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse(orderResult))
}

func GetOrderEating(c echo.Context) error {
	orderResult := repository.GetOrderStatustData(2)
	if orderResult == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve order eating data"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse(orderResult))
}
func GetOrderPaid(c echo.Context) error {
	orderResult := repository.GetOrderStatustData(4)
	if orderResult == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve order paid data"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse(orderResult))
}
func GetOrderCancel(c echo.Context) error {
	orderResult := repository.GetOrderStatustData(3)
	if orderResult == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve order cancel data"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse(orderResult))
}

func GetOrderDateHistory(c echo.Context) error {
	dateStr := c.Param("date")
	if dateStr == "" {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid date"))
	}
	orderResult := repository.GetOrderDateData(dateStr)
	if orderResult == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve order date history data"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse(orderResult))
}

func GetEmployeeIdOrder(c echo.Context) error {
	employeeId := c.Param("employeeId")
	if employeeId == "" {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid date"))
	}
	orderResult := repository.GetOrderEmployeeData(employeeId)
	if orderResult == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve employee order data"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse(orderResult))
}

func EatingOrder(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid order ID"))
	}

	err = repository.UpdateStatusOrderData(orderId, 2, "")
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, response2.ErrorResponse("Order not found"))
		}
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to Eating update to order"))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse("Order Update status Eating successfully"))
}

func CreateOrder(c echo.Context) error {
	var createOrder request.CreateOrderRequest
	if err := c.Bind(&createOrder); err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid request body"))
	}
	customerId, err := repository.CreateCustomer(createOrder.Number)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to create customer"))
	}
	tableId, err := repository.FindAvailableTable(createOrder.Number)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to find available table"))
	}
	employeeId, err := repository.FindEmployeeId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to find employee"))
	}
	orderId, err := repository.CreateOrder(tableId, customerId, employeeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to create order"))
	}
	for _, item := range createOrder.OrderList {
		err = repository.CreateOrderListItem(item.FoodId, item.Quantity, orderId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to create order list"))
		}
	}
	err = repository.UpdateStatusCreateOrder(tableId, employeeId, createOrder.Number)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to update order status"))
	}
	err = repository.UpdateTotalAmountOrder()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to update order total amount"))
	}
	return c.JSON(http.StatusOK, response2.SuccessResponse(map[string]interface{}{
		"customer_id": customerId,
		"table_id":    tableId,
		"employee_id": employeeId,
		"order_id":    orderId,
	}))
}
