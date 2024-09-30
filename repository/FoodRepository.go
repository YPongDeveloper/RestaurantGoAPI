package repository

import (
	"log"
	"restaurant/database"
	"restaurant/model/dao"
	"restaurant/model/response"
)

func GetMenuData() []dao.Food {
	var foods []dao.Food
	result := database.DB.Find(&foods)
	if result.Error != nil {
		return nil
	}
	return foods
}

func GetMenuIdData(foodId int) response.FoodMenuIdResponse {
	var foodMenuId response.FoodMenuIdResponse
	db := database.DB

	result := db.Raw("SELECT  f.food_id,f.food_name,c.category_name,f.description, f.price,f.available,f.calorie"+
		" FROM food f INNER JOIN "+
		"category c ON f.category_id = c.category_id where f.food_id = ?;", foodId).Scan(&foodMenuId)
	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return response.FoodMenuIdResponse{}
	}
	return foodMenuId
}

func GetMenuCategoryData(categoryId int) []response.FoodMenuIdResponse {
	var foodMenuCategory []response.FoodMenuIdResponse
	db := database.DB

	result := db.Raw("SELECT  f.food_id,f.food_name,c.category_name,f.description, f.price,f.available,f.calorie"+
		" FROM food f "+
		"INNER JOIN "+
		"category c ON f.category_id = c.category_id where c.category_id = ?;", categoryId).Scan(&foodMenuCategory)
	if result.Error != nil {
		log.Println("Error retrieving customer orders:", result.Error)
		return nil
	}
	return foodMenuCategory
}

func GetCategoryData() []response.CategoryResponse {
	var category []response.CategoryResponse
	db := database.DB

	result := db.Raw("select * from category;").Scan(&category)
	if result.Error != nil {
		log.Println("Error retrieving Categorys Data:", result.Error)
		return nil
	}
	return category
}
