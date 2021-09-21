package route

import (
	"myproperty-api/controllers"

	"github.com/labstack/echo"
)

func NewUser(app *echo.Echo) {
	app.GET("/users", controllers.GetAllUsersController)
	app.POST("/users", controllers.CreateUserController)
	app.GET("/users/:id", controllers.GetUserByIDController)
	app.DELETE("/users/:id", controllers.DeleteUserByIDController)
	app.PUT("/users/:id", controllers.UpdateUserByIDController)
}
