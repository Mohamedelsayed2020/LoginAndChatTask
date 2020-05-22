package main

import (
	"LoginAndChatTask/api/routes"
	"LoginAndChatTask/command/migration"
	"LoginAndChatTask/model/common"
)

func main() {
	migration.Execute()
	routes.Routes()
	go common.HandleMessages()
}