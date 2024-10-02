package repository

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"restaurant/database"
	"restaurant/model/response"
)

func UpdateStatusOrderData(orderId int, status int, review string) error {
	db := database.DB

	result := db.Exec("UPDATE Orders SET status = ?, review = ? WHERE order_id = ?", status, review, orderId)
	if result.Error != nil {
		log.Println("Error updating order status:", result.Error)
		return result.Error
	}
	if status == 4 || status == 3 {
		var tableId int
		result := db.Raw("SELECT table_id FROM Orders WHERE order_id = ?", orderId).Scan(&tableId)
		if result.Error != nil {
			log.Println("Error retrieving table_id:", result.Error)
			return result.Error
		}
		if tableId == 0 {
			log.Println("No table_id found for the given order_id")
			return errors.New("no table_id found for the given order_id")
		}
		updateResult := db.Exec("UPDATE Tables SET status = 0 WHERE table_id = ?", tableId)
		if updateResult.Error != nil {
			log.Println("Error updating table status:", updateResult.Error)
			return updateResult.Error
		}
		updateEmployeeQuery := "UPDATE employee SET status = 0 WHERE employee_id = ?"
		var employeeId int
		resultEMID := db.Raw("SELECT employee_id FROM Orders WHERE order_id = ?;", orderId).Scan(&employeeId)
		if resultEMID.Error != nil {
			log.Println("Error retrieving employee_id:", resultEMID.Error)
			return resultEMID.Error
		}
		if employeeId == 0 {
			log.Println("No employee_id found for the given order_id")
			return errors.New("no employee_id found for the given order_id")
		}
		resultEP := database.DB.Exec(updateEmployeeQuery, employeeId)

		if resultEP.Error != nil {
			log.Println("Failed to update employee status:", resultEP.Error)
			return resultEP.Error
		}

		if resultEP.RowsAffected != 0 {
			log.Println("No table found with the specified employee_id")
			return errors.New("no table found with the specified employee_id")
		}
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func GetOrderByIdData(orderId int) response.OrderIdResponse {
	var orderIds response.OrderIdResponse
	db := database.DB

	result := db.Raw("SELECT o.customer_id, o.order_id,o.employee_id, o.table_id, o.total_amount, o.order_date, o.status, o.review "+
		"FROM Orders o inner join order_list ol on o.order_id = ol.order_id where o.order_id = ? group by o.order_id;", orderId).
		Scan(&orderIds)

	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return response.OrderIdResponse{}
	}

	return orderIds

}

func GetOrderListOrderData(orderId int) []response.OrderListResponse {
	var orderList []response.OrderListResponse
	db := database.DB

	result := db.Raw("SELECT f.food_name,ol.order_id, cg.category_name, ol.quantity, f.price * ol.quantity AS total_price "+
		"FROM customer c "+
		"INNER JOIN Orders o ON c.customer_id = o.customer_id "+
		"INNER JOIN order_list ol ON o.order_id = ol.order_id "+
		"INNER JOIN food f ON ol.food_id = f.food_id "+
		"INNER JOIN category cg ON f.category_id = cg.category_id "+
		"WHERE o.order_id = ?;", orderId).Scan(&orderList)
	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return nil
	}
	return orderList
}

func GetOrderData() []response.OrderAllResponse {
	var order []response.OrderAllResponse
	db := database.DB

	result := db.Raw("select o.order_id,table_id,order_date,total_amount,status,customer_id,employee_id,review,count(o.order_id) as total_menu " +
		"FROM Orders o inner join order_list ol on o.order_id = ol.order_id  group by o.order_id;").
		Scan(&order)

	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return nil
	}

	return order
}

func GetOrderStatustData(status int) []response.OrderAllResponse {
	var order []response.OrderAllResponse
	db := database.DB

	result := db.Raw("select o.order_id,table_id,order_date,total_amount,status,customer_id,employee_id,review,count(o.order_id) as total_menu "+
		"FROM Orders o inner join order_list ol on o.order_id = ol.order_id where status = ? group by o.order_id;", status).
		Scan(&order)

	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return nil
	}

	return order
}

func GetOrderDateData(dateStr string) []response.OrderAllResponse {
	var order []response.OrderAllResponse
	db := database.DB

	result := db.Raw("select o.order_id,table_id,order_date,total_amount,status,customer_id,employee_id,review,count(o.order_id) as total_menu "+
		"FROM Orders o inner join order_list ol on o.order_id = ol.order_id where DATE(order_date) = ? group by o.order_id;", dateStr).
		Scan(&order)

	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return nil
	}

	return order
}

func GetOrderEmployeeData(employeeId string) []response.OrderAllResponse {
	var order []response.OrderAllResponse
	db := database.DB

	result := db.Raw("select o.order_id,table_id,order_date,total_amount,status,customer_id,employee_id,review,count(o.order_id) as total_menu "+
		"FROM Orders o inner join order_list ol on o.order_id = ol.order_id where employee_id = ? group by o.order_id;", employeeId).
		Scan(&order)

	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return nil
	}

	return order
}

func CreateCustomer(number int) (int, error) {
	query := "INSERT INTO customer (number) VALUES (?);"

	result := database.DB.Exec(query, number)
	if result.Error != nil {
		log.Println("Failed to create customer:", result.Error)
		return 0, result.Error
	}

	var customerId int
	err := database.DB.Raw("SELECT LAST_INSERT_ID()").Scan(&customerId).Error
	if err != nil {
		log.Println("Failed to retrieve customer ID:", err)
		return 0, err
	}

	return customerId, nil
}

func FindAvailableTable(chairNumber int) (int, error) {
	var tableId int
	query := `
    WITH AvailableTables AS (
        SELECT table_id, chair_number
        FROM Tables
        WHERE chair_number >= ? AND status = 0
    ),
    MinChairTable AS (
        SELECT table_id
        FROM AvailableTables
        WHERE chair_number = (SELECT MIN(chair_number) FROM AvailableTables)
    )
    SELECT MIN(table_id)
    FROM MinChairTable;
    `

	result := database.DB.Raw(query, chairNumber).Scan(&tableId)
	if result.Error != nil {
		log.Println("Failed to find available table:", result.Error)
		return 0, result.Error
	}

	return tableId, nil
}

func FindEmployeeId() (int, error) {
	var employeeId int
	query := `
    SELECT employee_id
    FROM employee
    WHERE status = 0
    ORDER BY RAND()
    LIMIT 1;
    `

	result := database.DB.Raw(query).Scan(&employeeId)
	if result.Error != nil {
		log.Println("Failed to find employee:", result.Error)
		return 0, result.Error
	}

	return employeeId, nil
}
func CreateOrder(tableId, customerId, employeeId int) (int, error) {
	query := `
    INSERT INTO Orders (table_id, order_date, total_amount, status, customer_id, employee_id, review) 
    VALUES (?, NOW(), 0, 0, ?, ?, "")
    `

	result := database.DB.Exec(query, tableId, customerId, employeeId)
	if result.Error != nil {
		log.Println("Failed to create order:", result.Error)
		return 0, result.Error
	}

	var orderId int
	result = database.DB.Raw("SELECT LAST_INSERT_ID()").Scan(&orderId)
	if result.Error != nil {
		log.Println("Failed to retrieve order ID:", result.Error)
		return 0, result.Error
	}

	return orderId, nil
}

func CreateOrderListItem(foodId, quantity, orderId int) error {
	query := `
    INSERT INTO order_list (quantity, order_id, food_id) 
    VALUES (?, ?, ?);
    `

	result := database.DB.Exec(query, quantity, orderId, foodId)
	if result.Error != nil {
		log.Println("Failed to create order list item:", result.Error)
		return result.Error
	}
	return nil
}

func UpdateStatusCreateOrder(tableId int, employeeId int) error {
	if tableId != 0 && employeeId != 0 {
		updateTableQuery := "UPDATE Tables SET status = 1 WHERE table_id = ?"

		result := database.DB.Exec(updateTableQuery, tableId)

		if result.Error != nil {
			log.Println("Failed to update table status:", result.Error)
			return result.Error
		}

		if result.RowsAffected == 0 {
			log.Println("No table found with the specified table_id")
			return errors.New("no table found with the specified table_id")
		}
		updateEmployeeQuery := "UPDATE employee SET status = 2 WHERE employee_id = ?"

		resultEP := database.DB.Exec(updateEmployeeQuery, employeeId)

		if resultEP.Error != nil {
			log.Println("Failed to update table status:", resultEP.Error)
			return resultEP.Error
		}

		if resultEP.RowsAffected == 0 {
			log.Println("No table found with the specified table_id")
			return errors.New("no table found with the specified table_id")
		}
	} else {
		log.Println("%d %d", tableId, employeeId)
	}
	return nil
}
