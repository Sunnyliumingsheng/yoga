package service

import (
	"api/db"
	"api/loger"
)

// 尝试service的另一种写法
// 新增一个课程
func (m *Message) InsertNewCourse(account string, recommendMaxNum, recommendMinNum int, courseName, courseSubject, introduction, introductionURL string, isGroup, isTeam, isVIP bool) {
	adminId, err := db.SelectAdminIdByAdminAccount(account)
	if err != nil {
		// 账号不存在
		m.IsSuccess = false
		m.Info = "账号不存在"
		m.HaveError = true
		return
	}
	err, isUnique := db.InsertNewCourse(adminId, recommendMaxNum, recommendMinNum, courseName, courseSubject, introduction, introductionURL, isGroup, isTeam, isVIP)
	if !isUnique {
		// 名字重复
		loger.Loger.Println("名字重复")
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	if err != nil {
		// 其他错误
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	m.HaveError = false
	m.Info = "成功新增一个课程"
	m.IsSuccess = true
}

// 删除一个课程
func (m *Message) DropCourseByName(name string) {
	err, isExist := db.DeleteCourseByName(name)
	if !isExist {
		// 课程不存在
		m.IsSuccess = false
		m.Info = "课程不存在"
		m.HaveError = false
		return
	}
	if err != nil {
		// 其他错误
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	m.HaveError = false
	m.Info = "成功删除课程"
	m.IsSuccess = true
}

// 检索所有的课程
func (m *Message) SelectCourse() {
	courses, err := db.SelectCourse()
	if err != nil {
		// 其他错误
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	m.Result = courses
	m.HaveError = false
	m.Info = "获取课程列表成功"
	m.IsSuccess = true
}

//本文件中这些对于message的操作方法很好,能批量复制
