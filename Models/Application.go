package Models

import (
	"github.com/EngMohamedFathy/golang-chat-app/Config"
	_ "github.com/go-sql-driver/mysql"
)

type Application struct {
	Id         uint   `json:",omitempty"`
	Token      string `json:"token"`
	Name       string `json:"name"`
	ChatsCount uint32 `json:"chats_count"`
	Chats      []Chat `json:",omitempty"`
}

func (app *Application) TableName() string {
	return "applications"
}

// CreateApplication ... Insert New Application data
func CreateApplication(application *Application) (err error) {
	if err = Config.DB.Create(application).Error; err != nil {
		return err
	}
	return nil
}

// GetApplication ... Fetch only one application by Token
func GetApplication(application *Application, token string) (err error) {
	if err = Config.DB.Where("token = ?", token).First(application).Error; err != nil {
		return err
	}
	return nil
}

// GetApplicationChats ... Fetch only one application chats by Token
func GetApplicationChats(application *Application, token string) (err error) {
	return nil
}
