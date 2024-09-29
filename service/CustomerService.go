// service/customerController.go
package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restaurant/model/response"
	"restaurant/repository"
)

func GetCustomers(c echo.Context) error {
	result := repository.GetCustomersData()
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve All customer data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}
