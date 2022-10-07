package Controllers

import (
	"fmt"
	"github.com/EngMohamedFathy/golang-chat-app/Config"
	helpers "github.com/EngMohamedFathy/golang-chat-app/Helpers"
	"github.com/EngMohamedFathy/golang-chat-app/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateMessageRequest struct {
	ApplicationToken string `json:"application_token"`
	ChatNumber       int    `json:"chat_number"`
	Body             string `json:"body"`
}

// CreateMessage ... Create Application Chat Message
func CreateMessage(c *gin.Context) {
	var err any = nil
	var chat Models.Chat
	var message Models.Message

	// to bind request
	var req CreateMessageRequest
	c.BindJSON(&req)

	// to get chat id from application token and chat number
	err = Models.GetChat(&chat, req.ApplicationToken, req.ChatNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.ApiResponse{Status: "Error", Message: "Chat Not Found", Data: err})
		return
	}

	// to get last message id
	var lastChatNumber NResult
	Config.DB.Raw("SELECT MAX(number) as n FROM messages WHERE chat_id = ?", chat.Id).Scan(&lastChatNumber)

	// to set chat values
	message.Number = lastChatNumber.N + 1
	message.Body = req.Body
	message.ChatId = chat.Id

	err = Models.CreateMessage(&message)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusExpectationFailed)
	}
	message.Id = 0
	message.ChatId = 0

	c.JSON(http.StatusCreated, helpers.ApiResponse{Status: "Success", Message: "Message Created", Data: message})
}
