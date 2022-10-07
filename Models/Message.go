package Models

import (
	"github.com/EngMohamedFathy/golang-chat-app/Config"
	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Id     uint   `json:",omitempty"`
	Number uint32 `json:"number"`
	ChatId uint   `json:"chat_id,omitempty"`
	Body   string `json:"body"`
}

func (app *Message) TableName() string {
	return "messages"
}

// CreateMessage ... Insert New Application data
func CreateMessage(message *Message) (err error) {
	if err = Config.DB.Create(message).Error; err != nil {
		return err
	}
	return nil
}

// GetMessage ... Fetch only one application chat message by Token
func GetMessage(message *Application, token string, chatNumber int) (err error) {
	return nil
}
