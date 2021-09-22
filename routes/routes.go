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
	e.POST("/developers", controllers.CreateDeveloperController)
	e.GET("/developers/:id", controllers.GetDeveloperByIDController)
	e.DELETE("/developers/:id", controllers.DeleteDeveloperByIDController)
	e.PUT("/developers/:id", controllers.UpdateDeveloperByIDController)
	e.GET("/chats", controllers.GetAllChatsController)
	e.POST("/chats", controllers.CreateChatController)
	e.GET("/chats/:id", controllers.GetChatByIDController)
	e.DELETE("/chats/:id", controllers.DeleteChatByIDController)
	e.PUT("/chats/:id", controllers.UpdateChatByIDController)
	e.GET("/properties", controllers.GetAllPropertiesController)
	e.POST("/properties", controllers.CreatePropertyController)
	e.GET("/properties/:id", controllers.GetPropertyByIDController)
	e.DELETE("/properties/:id", controllers.DeletePropertyByIDController)
	e.PUT("/properties/:id", controllers.UpdatePropertyByIDController)
	e.GET("/transactions", controllers.GetAllTransactionsController)
	e.POST("/transactions", controllers.CreateTransactionController)
	e.GET("/transactions/:id", controllers.GetTransactionByIDController)
	e.DELETE("/transactions/:id", controllers.DeleteTransactionByIDController)
	e.PUT("/transactions/:id", controllers.UpdateTransactionByIDController)
	e.GET("/wishlist", controllers.GetAllWishlistsController)
	e.POST("/wishlist", controllers.CreateWishlistController)
	e.GET("/wishlist/:id", controllers.GetWishlistByIDController)
	e.DELETE("/wishlist/:id", controllers.DeleteWishlistByIDController)
	e.PUT("/wishlist/:id", controllers.UpdateWishlistByIDController)
	e.GET("/property-type", controllers.GetAllPropertyTypesController)
	e.POST("/property-type", controllers.CreatePropertyTypeController)
	e.GET("/property-type/:id", controllers.GetPropertyTypeByIDController)
	e.DELETE("/property-type/:id", controllers.DeletePropertyTypeByIDController)
	e.PUT("/property-type/:id", controllers.UpdatePropertyTypeByIDController)
	return e
}
