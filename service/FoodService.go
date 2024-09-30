package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restaurant/model/response"
	"restaurant/repository"
	"strconv"
)

func GetMenu(c echo.Context) error {
	result := repository.GetMenuData()
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve All Menu data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}

func GetMenuById(c echo.Context) error {
	foodId, err := strconv.Atoi(c.Param("foodId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Customer Id"))
	}
	result := repository.GetMenuIdData(foodId)
	if result == (response.FoodMenuIdResponse{}) {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve All Menu data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}

func GetMenuByCategory(c echo.Context) error {
	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid Category Id"))
	}
	result := repository.GetMenuCategoryData(categoryId)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve All Menu data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}

func GetCategory(c echo.Context) error {
	result := repository.GetCategoryData()
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve All Menu data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}
