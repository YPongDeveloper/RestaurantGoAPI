package request

type createOrder struct {
	number      int    `json:"number"`
	table_id    int    `json:"table_id"`
	orderDate   string `json:"order_date"`
	totalAmount int    `json:"total_amount"`
	status      int    `json:"status"`
	customerId  int    `json:"customer_id"`
	employeeId  int    `json:"employee_id"`
}
