package controllers

import (
	"myproperty-api/lib/database"
	model "myproperty-api/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllChatsController(c echo.Context) error {
	chats := database.GetChats()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllchatsController",
		"data":    chats,
	})
}

func GetChatByIDController(c echo.Context) error {
	id := c.Param("id")
	chat := database.GetChatByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetChatByIDController",
		"data":    chat,
	})
}

func DeleteChatByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteChatByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeletechatByIDController",
	})
}

func UpdateChatByIDController(c echo.Context) error {
	id := c.Param("id")

	var chat model.Chat
	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateChatController",
			"error":   err.Error(),
		})
	}
	database.UpdateChatByID(id, chat)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetChatByIDController",
		"data":    chat,
	})
}

func CreateChatController(c echo.Context) error {
	var newchat model.Chat
	if err := c.Bind(&newchat); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatechatController",
			"error":   err.Error(),
		})
	}

	newchat = database.CreateChat(newchat)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreatechatController",
		"data":    newchat,
	})
}
