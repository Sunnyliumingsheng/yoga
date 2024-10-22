package cmd

import (
	"fmt"
	"time"
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
	fmt.Println("请您确定对每一个操作的危害都清楚了")
	time.Sleep(500 * time.Millisecond)

}
