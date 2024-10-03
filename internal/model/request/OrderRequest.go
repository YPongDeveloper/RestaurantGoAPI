package request

type ReviewRequest struct {
	Review string `json:"review"`
}
type CreateOrderRequest struct {
	Number    int                `json:"number"`
	OrderList []OrderListRequest `json:"order_list"`
}
type OrderListRequest struct {
	FoodId   int `json:"food_id"`
	Quantity int `json:"quantity"`
}
