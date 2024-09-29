package repository

import (
	"restaurant/database"
	"restaurant/model/dao"
)

func GetMenuData() []dao.Food {
	var foods []dao.Food
	result := database.DB.Find(&foods)
	if result.Error != nil {
		return nil
	}
	return foods
}
