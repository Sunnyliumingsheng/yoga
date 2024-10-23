package handler

import (
	"fmt"
	"os"

	"cli/cmd"
)

func HandleFlags() {
	if len(os.Args) < 2 {
		cmd.RootCmd()
		return
	}
	switch os.Args[1] {
	case "help":
		cmd.HelpCmd()
	case "config":
		cmd.ConfigCmd()
	case "login":
		cmd.LoginCmd()
	case "register":
		cmd.RegisterCmd()
	case "select_user_by_name":
		cmd.SelectUserByName()
	case "insert_user_by_name":
		cmd.InsertUserByName()
	case "drop_user_by_userId":
		cmd.DropUserByUserId()
	case "upgrade_user_level":
		cmd.UpgradeUserLevel()

	default:
		fmt.Println("未找到该命令，请检查并重新输入")
	}
}
