package cmd

import (
	"flag"
	"fmt"
	"os"
)

// 这里是根目录

func LoginCmd() {

	login := flag.NewFlagSet("login", flag.ExitOnError)
	var account string
	var password string

	login.StringVar(&account, "account", "", "输入您的账号,第一次使用的时候请先登录")
	login.StringVar(&password, "password", "", "输入密码,这两个都会明文存储在配置里")

	login.Parse(os.Args[2:])

	
}
func UpgradeUserLevel() {

}
