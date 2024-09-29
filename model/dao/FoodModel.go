package dao

type Food struct {
	FoodId      int    `json:"food_id"`
	FoodName    string `json:"food_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Available   int    `json:"available"`
	Calories    int    `json:"calories"`
	CategoryId  int    `json:"category_id"`
}

func (Food) TableName() string {
	return "food"
}
