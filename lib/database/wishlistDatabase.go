package database

import (
	"myproperty-api/config"
	model "myproperty-api/models"
)

func GetWishlists() []model.Wishlist {
	var wishlists []model.Wishlist
	config.DB.Preload("User", "Property").Find(&wishlists)
	return wishlists
}

func GetWishlistByID(id string) model.Wishlist {
	var wishlist model.Wishlist
	config.DB.Where("id = ?", id).Preload("User", "Property").Find(&wishlist)
	return wishlist
}

func CreateWishlist(wishlist model.Wishlist) model.Wishlist {
	config.DB.Create(&wishlist)
	return wishlist
}

func DeleteWishlistByID(id string) {
	var wishlist model.Wishlist
	config.DB.Where("id = ?", id).Delete(&wishlist)
}

func UpdateWishlistByID(id string, wishlist model.Wishlist) {
	config.DB.Where("id = ?", id).Updates(&wishlist)
}
