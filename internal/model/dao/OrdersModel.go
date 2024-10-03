package dao

type Orders struct {
	order_id    int    `json:"order_id"`
	table_id    int    `json:"table_id"`
	orderDate   string `json:"order_date"`
	totalAmount int    `json:"total_amount"`
	status      int    `json:"status"`
	customerId  int    `json:"customer_id"`
	employeeId  int    `json:"employee_id"`
	review      string `json:"review"`
}

func (Orders) TableName() string {
	return "Orders"
}
