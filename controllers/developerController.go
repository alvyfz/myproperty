package controllers

import (
	"myproperty-api/lib/database"
	model "myproperty-api/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllDeveloperController(c echo.Context) error {
	developers := database.GetDevelopers()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllDevelopersController",
		"data":    developers,
	})
}

func GetDeveloperByIDController(c echo.Context) error {
	id := c.Param("id")
	developer := database.GetDeveloperByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetDeveloperByIDController",
		"data":    developer,
	})
}

func DeleteDeveloperByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteDeveloperByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteDeveloperByIDController",
		"status":  "Success",
	})
}

func UpdateDeveloperByIDController(c echo.Context) error {
	id := c.Param("id")

	var developer model.Developer
	if err := c.Bind(&developer); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateDeveloperController",
			"error":   err.Error(),
		})
	}
	database.UpdateDeveloperByID(id, developer)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetDeveloperByIDController",
		"data":    developer,
	})
}

func CreateDeveloperController(c echo.Context) error {
	var newDeveloper model.Developer
	if err := c.Bind(&newDeveloper); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateDeveloperController",
			"error":   err.Error(),
		})
	}

	newDeveloper = database.CreateDeveloper(newDeveloper)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateUserController",
		"data":    newDeveloper,
	})
}
