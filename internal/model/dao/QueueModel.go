package dao

type Queue struct {
	QueueID int `gorm:"column:queue_id;primaryKey"`
	OrderID int `gorm:"column:order_id"`
	Status  int `gorm:"column:status"`
}
type QueueInfo struct {
	Number  int `json:"number"`   // จากตาราง customer
	QueueID int `json:"queue_id"` // จากตาราง queue
	OrderID int `json:"order_id"` // จากตาราง Orders
}

func (Queue) TableName() string {
	return "queue"
}
