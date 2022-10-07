package Controllers

import (
	"fmt"
	"github.com/EngMohamedFathy/golang-chat-app/Config"
	helpers "github.com/EngMohamedFathy/golang-chat-app/Helpers"
	"github.com/EngMohamedFathy/golang-chat-app/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NResult struct {
	N uint32 //or int ,or some else
}

type CreateRequest struct {
	ApplicationToken string `json:"application_token"`
}

// CreateChat ... Create Application Chat
func CreateChat(c *gin.Context) {
	var err any = nil
	var chat Models.Chat
	var application Models.Application

	// to bind request
	var req CreateRequest
	c.BindJSON(&req)

	// to get application id from application token
	err = Models.GetApplication(&application, req.ApplicationToken)
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.ApiResponse{Status: "Error", Message: "Application Not Found", Data: nil})
		return
	}

	// to get last message id
	var lastChatNumber NResult
	Config.DB.Raw("SELECT MAX(number) as n FROM chats WHERE application_id = ?", application.Id).Scan(&lastChatNumber)

	// to set chat values
	chat.Number = lastChatNumber.N + 1
	chat.ApplicationId = application.Id

	err = Models.CreateChat(&chat)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusExpectationFailed)
	}
	chat.Id = 0
	chat.ApplicationId = 0

	c.JSON(http.StatusCreated, helpers.ApiResponse{Status: "Success", Message: "Chat Created", Data: chat})
}
