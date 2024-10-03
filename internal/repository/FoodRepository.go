package repository

import (
	"errors"
	"log"
	"restaurant/database"
	"restaurant/internal/model/dao"
	"restaurant/internal/model/response"
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

func CreateFoodData(foodCreate dao.Food) error {
	db := database.DB

	query := "INSERT INTO food (food_name, description, price, available, calorie, category_id) VALUES (?, ?, ?, 0, ?, ?);"
	result := db.Exec(query, foodCreate.FoodName, foodCreate.Description, foodCreate.Price, foodCreate.Calorie, foodCreate.CategoryId)

	if result.Error != nil {
		log.Println("Error creating Food Data with raw SQL:", result.Error)
		return result.Error
	}

	log.Println("Food Data created successfully")
	return nil
}

func CreateCategoryData(createCategory dao.Category) error {
	db := database.DB

	query := "INSERT INTO category (category_name) VALUES (?);"
	result := db.Exec(query, createCategory.CategoryName)

	if result.Error != nil {
		log.Println("Error creating Category Data with raw SQL:", result.Error)
		return result.Error
	}

	log.Println("Category Data created successfully")
	return nil
}

func UpdateFoodData(food dao.Food, foodId int) error {
	db := database.DB

	query := "UPDATE food SET "
	params := []interface{}{}

	if food.FoodName != "" {
		query += "food_name = ?, "
		params = append(params, food.FoodName)
	}

	if food.CategoryId != 0 {
		query += "category_id = ?, "
		params = append(params, food.CategoryId)
	}

	if food.Description != "" {
		query += "description = ?, "
		params = append(params, food.Description)
	}

	if food.Price != 0 {
		query += "price = ?, "
		params = append(params, food.Price)
	}

	if food.Available != nil {
		query += "available = ?, "
		params = append(params, food.Available)
	}

	if food.Calorie != 0 {
		query += "calorie = ?, "
		params = append(params, food.Calorie)
	}

	if len(params) == 0 {
		log.Println("No fields to update")
		return errors.New("no fields to update")
	}

	query = query[:len(query)-2]

	query += " WHERE food_id = ?"
	params = append(params, foodId)

	result := db.Exec(query, params...)
	if result.Error != nil {
		log.Println("Error updating Food Data with raw SQL:", result.Error)
		return result.Error
	}

	log.Println("Food Data updated successfully")
	return nil
}
