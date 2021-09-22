package controllers

import (
	"myproperty-api/lib/database"
	model "myproperty-api/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllPropertyTypesController(c echo.Context) error {
	propertyTypes := database.GetPropertyTypes()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllPropertyTypesController",
		"data":    propertyTypes,
	})
}

func GetPropertyTypeByIDController(c echo.Context) error {
	id := c.Param("id")
	propertyType := database.GetPropertyTypeByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetPropertyTypeByIDController",
		"data":    propertyType,
	})
}

func DeletePropertyTypeByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeletePropertyTypeByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeletePropertyTypeByIDController",
	})
}

func UpdatePropertyTypeByIDController(c echo.Context) error {
	id := c.Param("id")

	var propertyType model.PropertyType
	if err := c.Bind(&propertyType); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatePropertyTypeController",
			"error":   err.Error(),
		})
	}
	database.UpdatePropertyTypeByID(id, propertyType)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetPropertyTypeByIDController",
		"data":    propertyType,
	})
}

func CreatePropertyTypeController(c echo.Context) error {
	var newPropertyType model.PropertyType
	if err := c.Bind(&newPropertyType); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatePropertyTypeController",
			"error":   err.Error(),
		})
	}

	newPropertyType = database.CreatePropertyType(newPropertyType)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreatePropertyTypeController",
		"data":    newPropertyType,
	})
}
