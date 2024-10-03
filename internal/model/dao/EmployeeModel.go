package dao

type Employee struct {
	EmployeeId int    `json:"employee_id" gorm:"column:employee_id"`
	FirstName  string `json:"first_name" gorm:"column:first_name"`
	LastName   string `json:"last_name" gorm:"column:last_name"`
	Status     int    `json:"status" gorm:"column:status"`
}

func (Employee) TableName() string {
	return "employee"
}
