package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"restaurant/internal/model/dao"
	response2 "restaurant/internal/model/response"
	"restaurant/internal/repository"
	"strconv"
)

func GetMenu(c echo.Context) error {
	result := repository.GetMenuData()
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve All Menu data"))
	}
	return c.JSON(http.StatusOK, response2.SuccessResponse(result))
}

func GetMenuById(c echo.Context) error {
	foodId, err := strconv.Atoi(c.Param("foodId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid Customer Id"))
	}
	result := repository.GetMenuIdData(foodId)
	if result == (response2.FoodMenuIdResponse{}) {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve Menu id data"))
	}
	return c.JSON(http.StatusOK, response2.SuccessResponse(result))
}

func GetMenuByCategory(c echo.Context) error {
	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid Category Id"))
	}
	result := repository.GetMenuCategoryData(categoryId)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve Menu category data"))
	}
	return c.JSON(http.StatusOK, response2.SuccessResponse(result))
}

func GetCategory(c echo.Context) error {
	result := repository.GetCategoryData()
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to retrieve All category data"))
	}
	return c.JSON(http.StatusOK, response2.SuccessResponse(result))
}

func CreateFood(c echo.Context) error {
	var foodCreate dao.Food
	if err := c.Bind(&foodCreate); err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid request body"))
	}
	repository.CreateFoodData(foodCreate)

	return c.JSON(http.StatusOK, response2.SuccessResponse("Create Food successfully"))
}

func CreateCategory(c echo.Context) error {
	var categoryCreate dao.Category
	if err := c.Bind(&categoryCreate); err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid request body"))
	}

	repository.CreateCategoryData(categoryCreate)

	return c.JSON(http.StatusOK, response2.SuccessResponse("Create category successfully"))
}

func EditFood(c echo.Context) error {
	foodId, err := strconv.Atoi(c.Param("foodId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid Food Id"))
	}
	var editFood dao.Food
	if err := c.Bind(&editFood); err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid request body"))
	}
	result := repository.UpdateFoodData(editFood, foodId)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse("Failed to update Food data"))
	}
	return c.JSON(http.StatusOK, response2.SuccessResponse("Update Food successfully"))
}

func AvailableFood(c echo.Context) error {
	foodId, err := strconv.Atoi(c.Param("foodId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid Food Id"))
	}
	status, err := strconv.Atoi(c.Param("status"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid Food Status"))
	}
	var updateFood dao.Food
	updateFood.Available = &status
	result := repository.UpdateFoodData(updateFood, foodId)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse(fmt.Sprintf("Failed to update Available Food data: %d", updateFood.Available)))
	}
	return c.JSON(http.StatusOK, response2.SuccessResponse("Available Food successfully"))
}

func DeleteFood(c echo.Context) error {

	foodId, err := strconv.Atoi(c.Param("foodId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response2.ErrorResponse("Invalid Food Id"))
	}

	status := 2
	var updateFood dao.Food
	updateFood.Available = &status

	result := repository.UpdateFoodData(updateFood, foodId)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, response2.ErrorResponse(fmt.Sprintf("Failed to Delete Food data")))
	}

	return c.JSON(http.StatusOK, response2.SuccessResponse("Delete Food successfully"))
}
