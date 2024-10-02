package repository

import (
	"log"
	"restaurant/database"
	"restaurant/model/dao"
)

func GetEmployees() []dao.Employee {
	var employees []dao.Employee

	result := database.DB.Find(&employees)
	if result.Error != nil {
		log.Println("Failed to retrieve data:", result.Error)
		return nil
	}
	return employees
}

func GetEmployeeByIdData(employeeId string) (*dao.Employee, error) {
	var employee dao.Employee

	result := database.DB.First(&employee, employeeId)
	if result.Error != nil {
		log.Println("Failed to retrieve employee data:", result.Error)
		return nil, result.Error
	}

	return &employee, nil
}

func CreateEmployeeData(create dao.Employee) {
	database.DB.Create(&create)
}

func UpdateEmployeeData(update dao.Employee, employeeId string) {
	database.DB.Model(&update).Where(employeeId).Updates(update)
}
func FireEmployeeData(employeeId string) error {
	db := database.DB
	query := "update employee set status = 3 where employee_id = ?;"
	result := db.Exec(query, employeeId)

	if result.Error != nil {
		log.Println("Failed to fire employee:", result.Error)
		return result.Error
	}
	return nil
}
func AdsentEmployeeData(employeeId string) error {
	db := database.DB
	query := "update employee set status = 2 where employee_id = ?;"
	result := db.Exec(query, employeeId)

	if result.Error != nil {
		log.Println("Failed to fire employee:", result.Error)
		return result.Error
	}
	return nil
}

func WorkEmployeeData(employeeId string) interface{} {
	db := database.DB
	query := "update employee set status = 0 where employee_id = ?;"
	result := db.Exec(query, employeeId)

	if result.Error != nil {
		log.Println("Failed to fire employee:", result.Error)
		return result.Error
	}
	return nil
}
