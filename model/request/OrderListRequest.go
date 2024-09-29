package request

type orderList struct {
	quantity float64 `json:"quantity"`
	orderId  int64   `json:"orderId"`
	foodId   int64   `json:"foodId"`
}
