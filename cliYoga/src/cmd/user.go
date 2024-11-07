package cmd

import (
	"flag"
	"fmt"
	"os"

	"cli/nets"
)

// 这里是根目录

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
func UpdateUserLevelByName() {
	updateByName := flag.NewFlagSet("update_user_level_by_name", flag.ExitOnError)
	var name string
	var wantStudent bool
	var wantTeacher bool
	var wantAdmin bool
	var level int
	updateByName.StringVar(&name, "name", "", "输入想要更改的用户的名字")
	updateByName.BoolVar(&wantStudent, "student", false, "如果想要赋予学生权限,请使用这个--student")
	updateByName.BoolVar(&wantTeacher, "teacher", false, "同上,不过是教师权限")
	updateByName.BoolVar(&wantAdmin, "admin", false, "管理员权限")
	updateByName.IntVar(&level, "level", -1, "这是另一种方式,而且是优先考虑这个,新用户为0,赋予学生身份则+1,赋予教师则+2,赋予管理员+4.比如想给某人管理员和学生,则--level 5")
	updateByName.Parse(os.Args[2:])
	if name == "" {
		fmt.Println("请一定输入正确的姓名,不能为空")
		return
	}
	if level != -1 {
		//确认了输入了level
		//首先清空直接指定的
		wantStudent = false
		wantTeacher = false
		wantAdmin = false
		// 然后根据level来重新赋值

		if level > 7 || level < 0 {
			fmt.Println("请正确输入数组0-7")
			return
		}
		if level >= 4 {
			fmt.Println("赋予管理员权限")
			wantAdmin = true
		}
		level = level % 4
		if level >= 2 {
			fmt.Println("赋予教师权限")
			wantTeacher = true
		}
		level = level % 2
		if level >= 1 {
			fmt.Println("赋予学生权限")
			wantStudent = true
		}
	}
	// 第一遍测试时没有出现任何问题,以后可以进行调整
	// fmt.Println(wantStudent)
	// fmt.Println(wantTeacher)
	// fmt.Println(wantAdmin)
	nets.UpdateUserLevel(name, wantStudent, wantTeacher, wantAdmin)
}
func SelectUserTail() {
	selectUser := flag.NewFlagSet("select_users", flag.ExitOnError)
	var tail int
	selectUser.IntVar(&tail, "num", 10, "输入你想要检索多少数量的用户")
	nets.SelectUserTail(tail)
}
func InsertTeacherAccountAndPassword() {
	insertTeacherAccount := flag.NewFlagSet("set_teacher_auth", flag.ExitOnError)
	var account string
	var password string
	var teacherId int
	insertTeacherAccount.IntVar(&teacherId, "teacher_id", -1, "输入老师的id,注意,请务必确定要输入正确的id")
	insertTeacherAccount.StringVar(&account, "account", "", "输入想要插入的账号")
	insertTeacherAccount.StringVar(&password, "password", "", "输入想要插入的密码")
	insertTeacherAccount.Parse(os.Args[2:])
	if teacherId == -1 || account == "" || password == "" {
		fmt.Println("错误，请输入所有参数再请求")
		return
	}
	nets.InsertTeacherAccountAndPassword(teacherId, account, password)
}
func InsertAdminAccountAndPassword() {
	insertAdminAccount := flag.NewFlagSet("set_admin_auth", flag.ExitOnError)
	var account string
	var adminId int
	var password string
	insertAdminAccount.IntVar(&adminId, "admin_id", -1, "输入管理员的id,注意,请务必确定要输入正确的id")
	insertAdminAccount.StringVar(&account, "account", "", "输入想要插入的账号")
	insertAdminAccount.StringVar(&password, "password", "", "输入想要插入的密码")
	insertAdminAccount.Parse(os.Args[2:])
	if adminId == -1 || password == "" || account == "" {
		fmt.Println("错误，请输入所有参数再请求")
		return
	}
	nets.InsertAdminAccountAndPassword(adminId, account, password)
}
func SelectTeacherInfoByName() {
	selectTeacher := flag.NewFlagSet("select_teacher_by_name", flag.ExitOnError)
	var name string
	selectTeacher.StringVar(&name, "name", "", "输入想要查询的教师的姓名")
	selectTeacher.Parse(os.Args[2:])
	if name == "" {
		fmt.Println("请一定输入正确的姓名,不能为空")
		return
	}
	nets.SelectTeacherInfoByName(name)
}
func SelectAdminInfoByName() {
	selectAdmin := flag.NewFlagSet("select_admin_by_name", flag.ExitOnError)
	var name string
	selectAdmin.StringVar(&name, "name", "", "输入想要查询的管理员的姓名")
	selectAdmin.Parse(os.Args[2:])
	if name == "" {
		fmt.Println("请一定输入正确的姓名,不能为空")
		return
	}
	nets.SelectAdminInfoByName(name)

}
