package main

import (
	"fmt"
	"github.com/EngMohamedFathy/golang-chat-app/Config"
	"github.com/EngMohamedFathy/golang-chat-app/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Database Status:", err)
	}

	defer Config.DB.Close()

	//Config.DB.AutoMigrate(&Models.Application{})

	r := Routes.SetupRouter()

	//running
	err = r.Run()
	if err != nil {
		fmt.Println("Server Status:", err)
	}
}
