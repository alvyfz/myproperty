package routes

import (
	"myproperty-api/constants"
	"myproperty-api/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//
	// User routes
	e.GET("/users", controllers.GetAllUsersController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:id", controllers.GetUserByIDController)
	e.DELETE("/users/:id", controllers.DeleteUserByIDController)
	e.PUT("/users/:id", controllers.UpdateUserByIDController)
	e.POST("/login", controllers.LoginUsersController)
	//
	// Developer routes
	e.GET("/developers", controllers.GetAllDeveloperController)
	e.POST("/developers", controllers.CreateDeveloperController)
	e.GET("/developers/:id", controllers.GetDeveloperByIDController)
	e.DELETE("/developers/:id", controllers.DeleteDeveloperByIDController)
	e.PUT("/developers/:id", controllers.UpdateDeveloperByIDController)
	//
	// Chats routes
	e.GET("/chats", controllers.GetAllChatsController)
	e.POST("/chats", controllers.CreateChatController)
	e.GET("/chats/:id", controllers.GetChatByIDController)
	e.DELETE("/chats/:id", controllers.DeleteChatByIDController)
	e.PUT("/chats/:id", controllers.UpdateChatByIDController)
	//
	// Properties routes
	e.GET("/properties", controllers.GetAllPropertiesController)
	e.POST("/properties", controllers.CreatePropertyController)
	e.GET("/properties/:id", controllers.GetPropertyByIDController)
	e.DELETE("/properties/:id", controllers.DeletePropertyByIDController)
	e.PUT("/properties/:id", controllers.UpdatePropertyByIDController)
	//
	// Transactions routes
	e.GET("/transactions", controllers.GetAllTransactionsController)
	e.POST("/transactions", controllers.CreateTransactionController)
	e.GET("/transactions/:id", controllers.GetTransactionByIDController)
	e.DELETE("/transactions/:id", controllers.DeleteTransactionByIDController)
	e.PUT("/transactions/:id", controllers.UpdateTransactionByIDController)
	//
	// Wishlists routes
	e.GET("/wishlist", controllers.GetAllWishlistsController)
	e.POST("/wishlist", controllers.CreateWishlistController)
	e.GET("/wishlist/:id", controllers.GetWishlistByIDController)
	e.DELETE("/wishlist/:id", controllers.DeleteWishlistByIDController)
	e.PUT("/wishlist/:id", controllers.UpdateWishlistByIDController)
	//
	// Property type routes
	e.GET("/property-type", controllers.GetAllPropertyTypesController)
	e.POST("/property-type", controllers.CreatePropertyTypeController)
	e.GET("/property-type/:id", controllers.GetPropertyTypeByIDController)
	e.DELETE("/property-type/:id", controllers.DeletePropertyTypeByIDController)
	e.PUT("/property-type/:id", controllers.UpdatePropertyTypeByIDController)
	//
	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetDetailControllers)

	return e
}
