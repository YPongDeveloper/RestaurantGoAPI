package dao

type Employee struct {
	employeeId   int    `json:"employeeId"`
	employeeName string `json:"employeeName"`
	status       string `json:"status"`
}

func (Employee) TableName() string {
	return "employee"
}
