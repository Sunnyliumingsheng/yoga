package service

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"api/db"
	"api/loger"
	"api/util"
	"api/weixin"
)

// 如果没有收到token或者session的时候调用,可能注册新用户,如果成功则返回一个db.Authentication,失败只会返回错误信息
func RegisterUser(code string) (message Message) {
	openid, err := weixin.GetOpenId(code)
	if err != nil {
		return Message{IsSuccess: false, HaveError: true, Info: "获取openid失败", Result: err.Error()}
	}
	isExist, userId, level, err := db.IsThisOpenIdExistedAndGetLevel(openid)
	if err != nil {
		loger.Loger.Println("!!!!!!!!!!!!严重错误, 在检查用户是否存在的时候遇到了除了openid不存在以外的错误", err.Error(), "openid:", openid)
		return Message{IsSuccess: false, HaveError: true, Info: "获取用户信息失败", Result: err.Error()}
	}
	if isExist {
		// 给出新的token和session给到用户
		tokenChan := make(chan string)
		go util.AsyncGenerateToken(string(userId), tokenChan)
		go db.AddSession(fmt.Sprint(userId), level)
		token := <-tokenChan
		return Message{IsSuccess: true, HaveError: false, Info: "获取token成功", Result: db.Authentication{Token: token, Session: fmt.Sprint(userId)}}
	} else {
		userId, err = db.InsertUserAndGetUserId(openid)
		if err != nil {
			loger.Loger.Println("!!!!!!!!!!!!严重错误, 检查了是否存在但仍然插入新用户失败", err.Error(), "openid:", openid)
			return Message{IsSuccess: false, HaveError: true, Info: "插入用户信息失败", Result: err.Error()}
		}
		tokenChan := make(chan string)
		go util.AsyncGenerateToken(string(userId), tokenChan)
		go db.AddSession(fmt.Sprint(userId), level)
		token := <-tokenChan
		return Message{
			IsSuccess: true,
			HaveError: false,
			Info:      "新用户注册成功",
			Result:    db.Authentication{Token: token, Session: fmt.Sprint(userId)},
		}
	}
}

// 如果客户端能提供session和token,成功就刷新session,失败就需要客户端重新调用另一个api
func SessionAndTokenAuthentication(session string, token string) (message Message) {
	userId, err := strconv.Atoi(session)
	if err != nil {
		return Message{IsSuccess: false, HaveError: true, Info: "session转换成int64失败", Result: err.Error()}
	}
	var isOnlineChan chan bool = make(chan bool)
	var levelChan chan int = make(chan int)
	var isValidChan chan bool = make(chan bool)
	go db.AsyncAuthSession(session, isOnlineChan, levelChan)
	go util.AsyncParseToken(token, isValidChan)
	//一般来说redis的速度远快于token计算
	level := <-levelChan
	isOnline := <-isOnlineChan
	if isOnline {
		// 为了防止访问着突然两个都恰好失效,这里需要重新给到session
		db.AddSession(session, level)
		return Message{IsSuccess: true, HaveError: false, Info: "session和验证通过,时间刷新", Result: nil}
	}
	isValid := <-isValidChan

	if isValid {
		level, err = db.IntUserIdSelectUserLevel(userId)
		if err != nil {
			loger.Loger.Println("!!!!!!!!!!!!严重错误, 在检查用户等级的时候遇到了err", err.Error(), "userId:", userId)
			return Message{IsSuccess: false, HaveError: true, Info: "获取用户等级失败", Result: err.Error()}
		}
		db.AddSession(session, level)
		return Message{IsSuccess: true, HaveError: false, Info: "token验证通过", Result: level}
	}
	return Message{IsSuccess: false, HaveError: false, Info: "session和token不匹配", Result: nil}

}

// 直接检索用户信息,注意这个是管理员和超级用户才有权限调用的,正常情况下是不可能让你知道别人的openid等信息的
func SelectUserInfoByName(name string) (message Message) {
	user, err := db.SelectUserInfoByName(name)
	if err == gorm.ErrRecordNotFound {
		return Message{IsSuccess: false, HaveError: false, Info: "用户不存在", Result: nil}
	}
	if err != nil {
		loger.Loger.Println("error: 通过name查询个人信息的时候出现了错误")
		return Message{IsSuccess: false, HaveError: true, Info: "查询用户信息失败", Result: nil}
	}
	return Message{IsSuccess: true, HaveError: false, Info: "查询用户信息成功", Result: user}
}
func InsertNewUser(name string) (message Message) {
	haveExist, err := db.InsertNewUser(name)
	if haveExist {
		return Message{IsSuccess: false, HaveError: false, Info: "该name已被占用", Result: nil}
	}
	if err != nil {
		loger.Loger.Println("error: 插入用户数据的时候失败 , name: ", name, "error : ", err)
		return Message{IsSuccess: false, HaveError: true, Info: "插入新用户数据失败", Result: nil}
	}
	return Message{IsSuccess: true, HaveError: false, Info: "插入新用户数据成功", Result: nil}
}
func DropUserByStringUserId(userId string) (message Message) {
	err := db.DropUserByStringUserId(userId)
	if err != nil {
		return Message{IsSuccess: false, HaveError: true, Info: "删除时出现错误" + err.Error(), Result: nil}
	}
	return Message{IsSuccess: true, HaveError: false, Info: "删除用户成功", Result: nil}
}
func UpdateUserLevel(name string, isStudent, isTeacher, isAdmin bool) (message Message) {
	user, err := db.SelectUserInfoByName(name)
	if err != nil {
		return Message{IsSuccess: false, HaveError: true, Info: "没有这个name的用户,请检查用户是否存在" + err.Error(), Result: nil}
	}

	//这里我本想追求效率,使用并发,可是实际上这段代码并不面向用户,而且对失误的容忍度很低,就用同步逻辑来做
	//这里都是屎山的一部分,本来想使用并发,但是又改了之后又不想改格式,只能这样凑合着用了.
	handle := make(chan error, 3)
	db.HandleUserLevelStudent(handle, &user, isStudent)
	db.HandleUserLevelTeacher(handle, &user, isTeacher)
	db.HandleUserLevelAdmin(handle, &user, isAdmin)
	for i := 0; i < 3; i++ {
		err = <-handle
		if err != nil {
			return Message{IsSuccess: false, HaveError: true, Info: "修改用户等级时遇到错误" + err.Error(), Result: nil}
		}
	}
	return Message{IsSuccess: true, HaveError: false, Info: "修改用户等级成功", Result: nil}

}
func Rename(userId string, newName string) (message Message) {
	err := db.Rename(userId, newName)
	if err != nil {
		return Message{IsSuccess: false, HaveError: true, Info: err.Error(), Result: nil}
	}
	return Message{IsSuccess: true, HaveError: false, Info: "修改name成功", Result: nil}
}

func (m *Message) SelectUserTail(tail int) {
	users, err := db.SelectUserTail(tail)
	if err != nil {
		m.IsSuccess = false
		m.HaveError = true
		m.Info = "查询用户列表时遇到错误" + err.Error()
		m.Result = nil
		return
	}
	m.IsSuccess = true
	m.HaveError = false
	m.Info = "查询用户列表成功"
	m.Result = users
}
func (m *Message) SelectAdminInfoByName(name string) {
	userInfo, isExist, err := db.SelectAndCheckUserInfoByName(name)
	if err != nil {
		m.IsSuccess = false
		m.HaveError = true
		m.Info = "查询这个用户信息时遇到错误" + err.Error()
		m.Result = nil
		loger.Loger.Println("error: ", "查询这个用户信息时遇到错误", err.Error())
		return
	}
	if !isExist {
		m.IsSuccess = false
		m.HaveError = false
		m.Info = "该用户不存在"
		m.Result = nil
		return
	}
	adminInfo, err := db.SelectAdminInfo(userInfo.UserID)
	if err != nil {
		m.IsSuccess = false
		m.HaveError = true
		m.Info = "查询这个管理员信息时遇到错误" + err.Error()
		m.Result = nil
		loger.Loger.Println("error: ", "查询这个管理员信息时遇到错误", err.Error())
		return
	}
	result := appendAdminInfo(adminInfo, userInfo)
	m.IsSuccess = true
	m.Result = result
	m.HaveError = false
	m.Info = "查询这个用户和管理员信息成功"
}
func (m *Message) SelectTeacherInfoByName(name string) {
	userInfo, isExist, err := db.SelectAndCheckUserInfoByName(name)
	if !isExist {
		m.IsSuccess = false
		m.HaveError = false
		m.Info = "该用户不存在"
		m.Result = nil
		return
	}
	if err != nil {
		m.IsSuccess = false
		m.HaveError = true
		m.Info = "查询这个用户信息时遇到错误" + err.Error()
		m.Result = nil
		loger.Loger.Println("error: ", "查询这个用户信息时遇到错误", err.Error())
		return
	}
	teacherInfo, err := db.SelectTeacherInfo(userInfo.UserID)
	if err != nil {
		m.IsSuccess = false
		m.HaveError = true
		m.Info = "查询这个教师信息时遇到错误" + err.Error()
		m.Result = nil
		loger.Loger.Println("error: ", "查询这个教师信息时遇到错误", err.Error())
		return
	}
	result := appendTeacherInfo(teacherInfo, userInfo)
	m.IsSuccess = true
	m.Result = result
	m.HaveError = false
	m.Info = "查询这个用户和教师信息成功"
}
func (m *Message) InsertAdminAccountAndPassword(adminId int, account string, password string) {
	isExist, err := db.InsertAdminAccountAndPassword(adminId, account, password)
	if isExist {
		m.IsSuccess = false
		m.HaveError = false
		m.Info = "该account已存在"
		m.Result = nil
		return
	}
	if err != nil {
		m.IsSuccess = false
		m.HaveError = true
		m.Info = "插入admin account and password时遇到错误" + err.Error()
		m.Result = nil
		loger.Loger.Println("error: ", "插入admin account and password时遇到错误", err.Error())
		return
	}
	m.IsSuccess = true
	m.HaveError = false
	m.Info = "插入成功"
}
func (m *Message) InsertTeacherAccountAndPassword(teacherId int, account, password string) {
	isExist, err := db.InsertTeacherAccountAndPassword(teacherId, account, password)
	if isExist {
		m.IsSuccess = false
		m.HaveError = false
		m.Info = "该account已存在"
		m.Result = nil
		return
	}
	if err != nil {
		m.IsSuccess = false
		m.HaveError = true
		m.Info = "插入teacher account and password时遇到错误" + err.Error()
		m.Result = nil
		loger.Loger.Println("error: ", "插入teacher account and password时遇到错误", err.Error())
		return
	}
	m.IsSuccess = true
	m.HaveError = false
	m.Info = "插入成功"
}
func (m *Message) AdminAndTeacherLogin(account, password string, level int) {
	if level == 1 {
		isOk, err := db.AdminLogin(account, password)
		if isOk {
			fmt.Println("登录成功")
			m.IsSuccess = true
			m.HaveError = false
			m.Info = "管理员登录成功"
			m.Result = util.GenerrateTokenWithLevelAndAccount(account, level)
			return
		} else {
			if err != nil {
				m.IsSuccess = false
				m.HaveError = true
				m.Info = "管理员登录时遇到错误" + err.Error()
				loger.Loger.Println("error:", "登录时遇到错误", err, "详细信息如下", account, password, level)
				return
			}
			fmt.Println("登录失败了")
			m.IsSuccess = false
			m.HaveError = false
			m.Info = "登录失败，账号不存在或者密码错误"
			return
		}

	}
	if level == 2 {
		isOk, err := db.TeacherLogin(account, password)
		if isOk {
			m.IsSuccess = true
			m.HaveError = false
			m.Info = "教师登录成功"
			m.Result = util.GenerrateTokenWithLevelAndAccount(account, level)
			return
		} else {
			if err != nil {
				m.IsSuccess = false
				m.HaveError = true
				m.Info = "教师登录时遇到错误" + err.Error()
				loger.Loger.Println("error:", "登录时遇到错误", err, "详细信息如下", account, password, level)
			}
			m.IsSuccess = false
			m.HaveError = false
			m.Info = "登录失败，账号不存在或者密码错误"
			return
		}

	}
	//前端必须要对level的值进行检查如果是别的值是不能进来的
	return
}
func (m *Message) SelectAdminOrTeacherInfoByAccount(account string, isAdmin bool) {
	if isAdmin {
		adminInfo, err := db.SelectAdminInfoByAccount(account)
		if err != nil {
			m.HaveError = true
			m.IsSuccess = false
			m.Info = "检索管理员失败" + err.Error()
			return
		}
		userInfo, err := db.SelectUserInfoByUserId(adminInfo.UserID)
		if err != nil {
			m.HaveError = true
			m.IsSuccess = false
			m.Info = "检索用户失败" + err.Error()
		}
		info := appendAdminInfo(adminInfo, userInfo)
		info.Password = ""
		m.HaveError = false
		m.IsSuccess = true
		m.Info = "检索成功"
		m.Result = info
	} else {
		teacherInfo, err := db.SelectTeacherInfoByAccount(account)
		if err != nil {
			m.HaveError = true
			m.IsSuccess = false
			m.Info = "检索教师信息失败" + err.Error()
			return
		}
		userInfo, err := db.SelectUserInfoByUserId(teacherInfo.UserID)
		if err != nil {
			m.HaveError = true
			m.IsSuccess = false
			m.Info = "检索用户失败" + err.Error()
		}
		info := appendTeacherInfo(teacherInfo, userInfo)
		info.Password = ""
		// 密码？ 还是不告诉好了吧
		m.HaveError = false
		m.IsSuccess = true
		m.Info = "检索成功"
		m.Result = info
	}
}
func (m *Message) UpdateUserInfo(userId string, nickname, signature string, gender bool) {
	err := db.UpdateUserInfo(userId, nickname, signature, gender)
	if err != nil {
		m.HaveError = true
		m.IsSuccess = false
		m.Info = err.Error()
	} else {
		m.HaveError = false
		m.IsSuccess = true
		m.Info = "更新成功"
	}
}
func (m *Message) UpdateTeacherInfo(userId string, introduction string) {
	err := db.UpdateTeacherInfo(userId, introduction)
	if err != nil {
		m.HaveError = true
		m.IsSuccess = false
		m.Info = err.Error()
	} else {
		m.HaveError = false
		m.IsSuccess = true
		m.Info = "更新成功"
	}
}
