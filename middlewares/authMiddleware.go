package middlewares

import "github.com/labstack/echo"

func BasicAuth(email, password string, c echo.Context) (bool, error) {
	if email == "admin@mail.com" && password == "admin" {
		return true, nil
	}
	return false, nil
}
