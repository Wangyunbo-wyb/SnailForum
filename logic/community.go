package logic

import (
	"SnailForum/config"
	"SnailForum/model"
)

func GetCommunityCategory() []model.Category {
	db := config.GetDB()
	category := make([]model.Category, 0)
	db.Where(&model.Category{}).Find(&category)
	return category
}

func GetCategoryById(cid int32) model.Category {
	db := config.GetDB()
	var category model.Category
	db.Where(cid).Find(&category)
	return category
}
