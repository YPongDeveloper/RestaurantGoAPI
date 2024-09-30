package repository

import (
	"log"
	"restaurant/database"
	"restaurant/model/dao"
	"restaurant/model/response"
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

func GetCustomerByIdData(customerId int) []response.CustomerIdResponse {
	var customerIds []response.CustomerIdResponse
	db := database.DB

	// Query ข้อมูลจากฐานข้อมูล
	result := db.Raw("SELECT c.customer_id, c.number, o.order_id, o.table_id, o.total_amount, o.order_date, o.status, o.review "+
		"FROM customer c INNER JOIN Orders o ON c.customer_id = o.customer_id WHERE c.customer_id = ?", customerId).
		Scan(&customerIds)

	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return nil
	}

	// ส่งคืนค่า customerIds
	return customerIds

}

func GetOrderListCustomerData(customerId int) []response.OrderListResponse {
	var orderList []response.OrderListResponse
	db := database.DB

	result := db.Raw("SELECT f.food_name,ol.order_id, cg.category_name, ol.quantity, f.price * ol.quantity AS total_price "+
		"FROM customer c "+
		"INNER JOIN Orders o ON c.customer_id = o.customer_id "+
		"INNER JOIN order_list ol ON o.order_id = ol.order_id "+
		"INNER JOIN food f ON ol.food_id = f.food_id "+
		"INNER JOIN category cg ON f.category_id = cg.category_id "+
		"WHERE c.customer_id = ?;", customerId).Scan(&orderList)
	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return nil
	}
	return orderList
}
