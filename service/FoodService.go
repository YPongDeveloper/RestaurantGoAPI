package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restaurant/model/response"
	"restaurant/repository"
)

func GetMenu(c echo.Context) error {
	result := repository.GetMenuData()
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve All Menu data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}
