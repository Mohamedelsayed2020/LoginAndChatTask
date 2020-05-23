package migration

import (
	"LoginAndChatTask/api/server"
	"LoginAndChatTask/core/common"
	"LoginAndChatTask/core/model"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var command = &cobra.Command{}

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "pg-migrate command is a sub command from root command",
	Long:  "pg-migrate command is a sub command from root command used to migrate postgres db changes",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {

		}
		db := server.Conn().Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
		if db.Error != nil {
			common.Logger("error=========","")
		}
		models := args[0]
		switch models {
		case "user":
			server.Conn().AutoMigrate(&model.User{})
		case "session":
			server.Conn().AutoMigrate(&model.Session{})
		case "message":
			server.Conn().AutoMigrate(&model.Message{})
		default:
			fmt.Println("This model hasn't created yet :(")
		}
		log.Println("Successfully Migrated :)")
	},
}

func init() {
	command.AddCommand(migrate)
}

func Execute() {
	if err := command.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
