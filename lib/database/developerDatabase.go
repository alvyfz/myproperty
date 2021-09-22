package database

import (
	"myproperty-api/config"
	model "myproperty-api/models"
)

func GetDevelopers() []model.Developer {
	var developers []model.Developer
	config.DB.Preload("User").Find(&developers)
	return developers
}

func GetDeveloperByID(id string) model.Developer {
	var developer model.Developer
	config.DB.Where("id = ?", id).Preload("User").Find(&developer)
	return developer
}

func CreateDeveloper(developer model.Developer) model.Developer {
	config.DB.Create(&developer)
	return developer
}

func DeleteDeveloperByID(id string) {
	var developer model.Developer
	config.DB.Where("id = ?", id).Delete(&developer)
}

func UpdateDeveloperByID(id string, developer model.Developer) {
	config.DB.Where("id = ?", id).Updates(&developer)
}
