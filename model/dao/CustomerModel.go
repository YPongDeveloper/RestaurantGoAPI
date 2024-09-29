package dao

type Customer struct {
	CustomerId int `json:"customerId" gorm:"column:customer_id"`
	Number     int `json:"number" gorm:"column:number"`
}

func (Customer) TableName() string {
	return "customer"
}

// Getter สำหรับ customerId
func (c *Customer) GetCustomerId() int {
	return c.CustomerId
}

// Setter สำหรับ customerId
func (c *Customer) SetCustomerId(id int) {
	c.CustomerId = id
}
func (c *Customer) GetNumber() int {
	return c.Number
}

// Setter สำหรับ customerId
func (c *Customer) SetNumber(id int) {
	c.Number = id
}
