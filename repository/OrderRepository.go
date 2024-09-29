package repository

import (
	"gorm.io/gorm"
	"log"
	"restaurant/database"
)

func UpdateStatusOrderData(orderId int, status int, review string) error {
	// เชื่อมต่อฐานข้อมูล
	db := database.DB
	if status == 4 || status == 3 {
		result := db.Exec("UPDATE Orders SET status = ?, review = ? WHERE order_id = ?", status, review, orderId)
		if result.Error != nil {
			log.Println("Error updating order status:", result.Error)
			return result.Error
		}

		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
	} else {
		return nil
	}

}
