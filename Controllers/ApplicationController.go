package Controllers

import (
	"fmt"
	helpers "github.com/EngMohamedFathy/golang-chat-app/Helpers"
	"github.com/EngMohamedFathy/golang-chat-app/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateApplication ... Create Application
func CreateApplication(c *gin.Context) {
	var err any = nil
	var application Models.Application

	// to set application post data in struct model
	c.BindJSON(&application)

	// to generate url safe token length 16
	application.Token, _ = helpers.UrlSafeBase64(16, true)

	// to create application
	err = Models.CreateApplication(&application)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusExpectationFailed)
		return
	}
	// to remove key from response
	application.Id = 0

	c.JSON(http.StatusCreated, helpers.ApiResponse{Status: "Success", Message: "Application Created", Data: application})
}
