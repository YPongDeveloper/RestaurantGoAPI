package repository

import (
	"log"
	"restaurant/database"
	"restaurant/model/dao"
)

func GetCustomersData() []dao.Customer {
	var customers []dao.Customer

	result := database.DB.Find(&customers)
	if result.Error != nil {
		log.Println("Failed to retrieve data:", result.Error)
		return nil
	}
	return customers
}
