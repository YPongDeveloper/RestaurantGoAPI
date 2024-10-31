package dao

type Employee struct {
	EmployeeId int    `json:"employee_id" gorm:"column:employee_id"`
	FirstName  string `json:"first_name" gorm:"column:first_name"`
	LastName   string `json:"last_name" gorm:"column:last_name"`
	Status     int    `json:"status" gorm:"column:status"`
	Position   int    `json:"position" gorm:"column:position"`
	Username   string `json:"username" gorm:"column:username"`
	Password   string `json:"password" gorm:"column:password"`
}

func (Employee) TableName() string {
	return "employee"
}
