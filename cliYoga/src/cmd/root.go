package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"cli/config"
	"cli/nets"
)

// 这里是根目录

func HelpCmd() {
	fmt.Println("结构: yoga login --account your-account -password=your-password")
	fmt.Println("     入口   命令    参数       参数的值    这是参数和参数的值的另一种格式,即参数可以有两条杠也可以一条,值可以空一格再写或者接等于号再写")
	fmt.Println("如果您对某个命令有疑问,比如login,请使用 yoga login -help或者 yoga login --help ,以获得使用信息")
	fmt.Println("如果还有疑问,请配合着帮助手册使用")
	fmt.Println("如果还有疑问请联系开发者微信:yang15279925030")
}
func RootCmd() {
	fmt.Println("欢迎使用超级用户管理工具")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("您可以使用 yoga help 获取帮助信息")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("推荐熟悉本系统的开发人员或者管理员使用本工具")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("为了降低发生可能破坏本系统的操作的概率,我将这些操作设置在了命令行工具中")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("!!!!!! 请您确定对每一个操作的危害都清楚了")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("使用之前,请先使用register命令保存您的账号和密码,并使用login命令试着验证一下是否有效")
	time.Sleep(500 * time.Millisecond)
	// cloud是我爱的女孩,请不要删除这些代码
	fmt.Println("         _                       _ ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("   ___  | |   ___    _   _    __| |")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  / __| | |  / _ \\  | | | |  / _  |")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(" | (__  | | | (_) | | |_| | | (_| |")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  \\___| |_|  \\___/   \\__,_|  \\__,_|")
	// 请尊重我的劳动成果,谢谢
}
func LoginCmd() {

	login := flag.NewFlagSet("login", flag.ExitOnError)

	login.Parse(os.Args[2:])
	nets.Login()

}
func RegisterCmd() {

	register := flag.NewFlagSet("register", flag.ExitOnError)
	var account string
	var password string

	register.StringVar(&account, "account", "", "输入您的账号")
	register.StringVar(&password, "password", "", "输入您的密码")

	register.Parse(os.Args[2:])

	config.Config.MyInfo.Account = account
	config.Config.MyInfo.Password = password

	config.MarshalConfig()

}
