package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restaurant/model/dao"
	"restaurant/model/response"
	"restaurant/repository"
)

func GetEmployees(c echo.Context) error {
	result := repository.GetEmployees()
	if result == nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve Employees all data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}

func GetEmployeeById(c echo.Context) error {
	employeeId := c.Param("employeeId")
	result, err := repository.GetEmployeeByIdData(employeeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve employee data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse(result))
}

func CreateEmployee(c echo.Context) error {
	var employeeCreate dao.Employee
	if err := c.Bind(&employeeCreate); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request body"))
	}
	repository.CreateEmployeeData(employeeCreate)

	return c.JSON(http.StatusOK, response.SuccessResponse("Create Employee successfully"))
}
func UpdateEmployee(c echo.Context) error {
	employeeId := c.Param("employeeId")
	var employeeUpdate dao.Employee
	if err := c.Bind(&employeeUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request body"))
	}
	repository.UpdateEmployeeData(employeeUpdate, employeeId)

	return c.JSON(http.StatusOK, response.SuccessResponse("Update Employee successfully"))
}

func FireEmployee(c echo.Context) error {
	employeeId := c.Param("employeeId")

	result := repository.FireEmployeeData(employeeId)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve Employee data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse("Fire Employee Success!"))

}

func AbsentEmployee(c echo.Context) error {
	employeeId := c.Param("employeeId")

	result := repository.AdsentEmployeeData(employeeId)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve Employee data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse("Adsent Employee Success!"))
}

func WorkEmployee(c echo.Context) error {
	employeeId := c.Param("employeeId")

	result := repository.WorkEmployeeData(employeeId)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve Employee data"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponse("Enter Work Employee Success!"))
}
