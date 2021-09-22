package database

import (
	"myproperty-api/config"
	model "myproperty-api/models"
)

func GetPropertyTypes() []model.PropertyType {
	var propertyTypes []model.PropertyType
	config.DB.Find(&propertyTypes)
	return propertyTypes
}

func GetPropertyTypeByID(id string) model.PropertyType {
	var propertyType model.PropertyType
	config.DB.Where("id = ?", id).Find(&propertyType)
	return propertyType
}

func CreatePropertyType(propertyType model.PropertyType) model.PropertyType {
	config.DB.Create(&propertyType)
	return propertyType
}

func DeletePropertyTypeByID(id string) {
	var propertyType model.PropertyType
	config.DB.Where("id = ?", id).Delete(&propertyType)
}

func UpdatePropertyTypeByID(id string, propertyType model.PropertyType) {
	config.DB.Where("id = ?", id).Updates(&propertyType)
}
