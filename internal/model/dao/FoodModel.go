package dao

type Food struct {
	FoodId      int    `json:"food_id"`
	FoodName    string `json:"food_name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Available   *int   `json:"available"`
	Calorie     int    `json:"calories"`
	CategoryId  int    `json:"category_id"`
}
type Category struct {
	CategoryName string `json:"category_name"`
	CategoryId   int    `json:"category_id"`
}

func (Food) TableName() string {
	return "food"
}
