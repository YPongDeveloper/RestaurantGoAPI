package response

import "time"

type OrderIdResponse struct {
	OrderId     int       `json:"order_id"`
	CustomerId  int       `json:"customer_id"`
	EmployeeId  int       `json:"employee_id"`
	TableId     int       `json:"table_id"`
	TotalAmount int       `json:"total_amount"`
	OrderDate   time.Time `json:"order_date"`
	Status      int       `json:"status"`
	Review      string    `json:"review"`
}
type OrderResponse struct {
	CustomerId  int                 `json:"customer_id"`
	EmployeeId  int                 `json:"employee_id"`
	OrderId     int                 `json:"order_id"`
	TableId     int                 `json:"table_id"`
	TotalAmount int                 `json:"total_amount"`
	OrderDate   time.Time           `json:"order_date"`
	Status      int                 `json:"status"`
	Review      string              `json:"review"`
	OrderList   []OrderListResponse `json:"order_list"` // รายการอาหาร
}
type OrderAllResponse struct {
	OrderId     int       `json:"order_id"`
	CustomerId  int       `json:"customer_id"`
	Number      int       `json:"number"`
	EmployeeId  int       `json:"employee_id"`
	TableId     int       `json:"table_id"`
	TotalMenu   int       `json:"total_menu"`
	TotalAmount int       `json:"total_amount"`
	OrderDate   time.Time `json:"order_date"`
	Status      int       `json:"status"`
	Review      string    `json:"review"`
}
