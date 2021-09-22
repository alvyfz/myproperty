package controllers

import (
	"myproperty-api/lib/database"
	model "myproperty-api/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllWishlistsController(c echo.Context) error {
	wishlists := database.GetWishlists()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllwishlistsController",
		"data":    wishlists,
	})
}

func GetWishlistByIDController(c echo.Context) error {
	id := c.Param("id")
	wishlist := database.GetWishlistByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetwishlistByIDController",
		"data":    wishlist,
	})
}

func DeleteWishlistByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteWishlistByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeletewishlistByIDController",
	})
}

func UpdateWishlistByIDController(c echo.Context) error {
	id := c.Param("id")

	var wishlist model.Wishlist
	if err := c.Bind(&wishlist); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatewishlistController",
			"error":   err.Error(),
		})
	}
	database.UpdateWishlistByID(id, wishlist)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetwishlistByIDController",
		"data":    wishlist,
	})
}

func CreateWishlistController(c echo.Context) error {
	var newWishlist model.Wishlist
	if err := c.Bind(&newWishlist); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatewishlistController",
			"error":   err.Error(),
		})
	}

	newWishlist = database.CreateWishlist(newWishlist)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreatewishlistController",
		"data":    newWishlist,
	})
}
