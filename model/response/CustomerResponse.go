package response

import "time"

type CustomerOrderResponse struct {
	CustomerId  int                 `json:"customer_id"`
	Number      int                 `json:"number"`
	OrderId     int                 `json:"order_id"`
	TableId     int                 `json:"table_id"`
	TotalAmount int                 `json:"total_amount"`
	OrderDate   time.Time           `json:"order_date"`
	Status      int                 `json:"status"`
	Review      string              `json:"review"`
	OrderList   []OrderListResponse `json:"order_list"` // รายการอาหาร
}
type CustomerIdResponse struct {
	CustomerId  int       `json:"customer_id"`
	Number      int       `json:"number"`
	OrderId     int       `json:"order_id"`
	TableId     int       `json:"table_id"`
	TotalAmount int       `json:"total_amount"`
	OrderDate   time.Time `json:"order_date"`
	Status      int       `json:"status"`
	Review      string    `json:"review"`
}
type OrderListResponse struct {
	OrderId      int    `json:"order_id"`
	FoodName     string `json:"food_name"`
	CategoryName string `json:"category_name"`
	Quantity     int    `json:"quantity"`
	TotalPrice   int    `json:"total_price"`
}
