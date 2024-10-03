package dao

type Customer struct {
	CustomerId int `json:"customerId" gorm:"column:customer_id"`
	Number     int `json:"number" gorm:"column:number"`
}

func (Customer) TableName() string {
	return "customer"
}
