package response

type FoodMenuIdResponse struct {
	FoodId       int    `json:"food_id"`
	FoodName     string `json:"food_name"`
	ImageBase64  string `json:"image_base64"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Available    int    `json:"available"`
	Calorie      int    `json:"calorie"`
}
type CategoryResponse struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
