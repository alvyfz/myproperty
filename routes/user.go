package routes

import (
	"myproperty-api/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetAllUsersController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:id", controllers.GetUserByIDController)
	e.DELETE("/users/:id", controllers.DeleteUserByIDController)
	e.PUT("/users/:id", controllers.UpdateUserByIDController)
	e.GET("/developers", controllers.GetAllDeveloperController)
	e.POST("/developers", controllers.CreateUserController)
	e.GET("/developers/:id", controllers.GetUserByIDController)
	e.DELETE("/developers/:id", controllers.DeleteUserByIDController)
	e.PUT("/developers/:id", controllers.UpdateUserByIDController)
	return e
}
