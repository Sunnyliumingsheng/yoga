package cmd

import (
	"flag"
	"os"

	"cli/nets"
)

// 这里是根目录

func UpgradeUserLevel() {

}
func SelectUserByName() {
	selectByName := flag.NewFlagSet("select_user_by_name", flag.ExitOnError)
	var name string
	selectByName.StringVar(&name, "name", "", "输入想要查询的名字,你将获得这个人的详细信息")

	selectByName.Parse(os.Args[2:])

	nets.SelectUserInfoByName(name)
}
func InsertUserByName() {
	insertByName := flag.NewFlagSet("insert_user_by_name", flag.ExitOnError)
	var name string
	insertByName.StringVar(&name, "name", "", "输入想要插入的名字, 这个人将被插入到数据库中")

	insertByName.Parse(os.Args[2:])

	nets.InsertUserByName(name)
}
func DropUserByUserId() {
	dropByName := flag.NewFlagSet("drop_user_by_userId", flag.ExitOnError)
	var userId string
	dropByName.StringVar(&userId, "userId", "", "输入想要删除的 userId, 这个人将被从数据库中删除")

	dropByName.Parse(os.Args[2:])

	nets.DropUserByUserId(userId)
}
