package Models

import (
	"github.com/EngMohamedFathy/golang-chat-app/Config"
	_ "github.com/go-sql-driver/mysql"
)

type Chat struct {
	Id            uint      `json:",omitempty"`
	Number        uint32    `json:"number"`
	ApplicationId uint      `json:"application_id,omitempty"`
	MessagesCount uint32    `json:"messages_count"`
	Messages      []Message `json:",omitempty"`
}

func (app *Chat) TableName() string {
	return "chats"
}

type NResult struct {
	N uint32 //or int ,or some else
}

// CreateChat ... Insert New Application data
func CreateChat(chat *Chat) (err error) {
	if err = Config.DB.Create(chat).Error; err != nil {
		return err
	}
	return nil
}

// GetChat ... Fetch only one application by Token
func GetChat(chat *Chat, token string, chatNumber int) (err error) {
	//fmt.Printf("SELECT chats.* FROM chats inner join applications on chats.application_id = applications.id WHERE token = %v AND number= %v", token, chatNumber)
	if err = Config.DB.Raw("SELECT chats.* FROM chats inner join applications on chats.application_id = applications.id WHERE token = ? AND number= ?", token, chatNumber).Scan(&chat).Error; err != nil {
		return err
	}
	return nil
}

// GetChatMessages ... Fetch only one application chat messages by Token
func GetChatMessages(chat *Chat, token string, chatNumber int) (err error) {
	return nil
}
