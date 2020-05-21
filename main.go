package main

import (
	"LoginAndChatTask/api/routes"
	"LoginAndChatTask/command/migration"
)

func main() {
	migration.Execute()
	routes.Routes()
}