package controllers

import (
	"myproperty-api/lib/database"
	model "myproperty-api/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllpropertiesController(c echo.Context) error {
	properties := database.GetProperties()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllpropertysController",
		"data":    properties,
	})
}

func GetPropertyByIDController(c echo.Context) error {
	id := c.Param("id")
	property := database.GetPropertyByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetpropertyByIDController",
		"data":    property,
	})
}

func DeletePropertyByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeletePropertyByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeletepropertyByIDController",
	})
}

func UpdatePropertyByIDController(c echo.Context) error {
	id := c.Param("id")

	var property model.Property
	if err := c.Bind(&property); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatepropertyController",
			"error":   err.Error(),
		})
	}
	database.UpdatePropertyByID(id, property)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetPropertyByIDController",
		"data":    property,
	})
}

func CreatePropertyController(c echo.Context) error {
	var newProperty model.Property
	if err := c.Bind(&newProperty); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatepropertyController",
			"error":   err.Error(),
		})
	}

	newProperty = database.CreateProperty(newProperty)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreatepropertyController",
		"data":    newProperty,
	})
}
