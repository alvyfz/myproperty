package middlewares

import (
	"myproperty-api/config"
	model "myproperty-api/models"

	"github.com/labstack/echo"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user model.User
	if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
