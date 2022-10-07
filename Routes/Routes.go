package Routes

import (
	"github.com/EngMohamedFathy/golang-chat-app/Controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	{
		// applications routes
		//apiV1.GET("applications/:token", Controllers.GetApplication)
		apiV1.POST("applications", Controllers.CreateApplication)

		// chats routes
		//apiV1.GET("applications/:token/chats", Controllers.GetChats)
		apiV1.POST("chats", Controllers.CreateChat)

		//messages routes
		apiV1.POST("messages", Controllers.CreateMessage)

	}
	return r
}
