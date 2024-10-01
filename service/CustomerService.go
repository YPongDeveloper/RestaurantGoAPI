package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restaurant/model/response"
	"restaurant/repository"
	"strconv"
)

func GetCustomers(c echo.Context) error {
	result := repository.GetCustomersData()
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve All customer data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}

func GetCustomerById(c echo.Context) error {
	customerId, err := strconv.Atoi(c.Param("customerId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Customer Id"))
	}

	customerResult := repository.GetCustomerByIdData(customerId)
	if customerResult == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve All customer data"))
	}
	orderListResult := repository.GetOrderListCustomerData(customerId)
	if orderListResult == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve order list data"))
	}
	var customerOrderResponses []response.CustomerOrderResponse

	for _, customer := range customerResult {
		var customerOrderResponse response.CustomerOrderResponse
		customerOrderResponse.CustomerId = customer.CustomerId
		customerOrderResponse.EmployeeId = customer.EmployeeId
		customerOrderResponse.Number = customer.Number
		customerOrderResponse.OrderId = customer.OrderId
		customerOrderResponse.TotalAmount = customer.TotalAmount
		customerOrderResponse.OrderDate = customer.OrderDate
		customerOrderResponse.Status = customer.Status
		customerOrderResponse.Review = customer.Review
		for _, orderList := range orderListResult {
			if orderList.OrderId == customer.OrderId {
				customerOrderResponse.OrderList = append(customerOrderResponse.OrderList, orderList)
			}
		}
		customerOrderResponses = append(customerOrderResponses, customerOrderResponse)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(customerOrderResponses))
}
