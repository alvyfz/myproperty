package database

import (
	"myproperty-api/config"
	model "myproperty-api/models"
)

func GetProperties() []model.Property {
	var properties []model.Property
	config.DB.Preload("PropertyType", "Developer").Find(&properties)
	return properties
}

func GetPropertyByID(id string) model.Property {
	var property model.Property
	config.DB.Where("id = ?", id).Preload("PropertyType", "Developer").Find(&property)
	return property
}

func CreateProperty(property model.Property) model.Property {
	config.DB.Create(&property)
	return property
}

func DeletePropertyByID(id string) {
	var property model.Property
	config.DB.Where("id = ?", id).Delete(&property)
}

func UpdatePropertyByID(id string, property model.Property) {
	config.DB.Where("id = ?", id).Updates(&property)
}
