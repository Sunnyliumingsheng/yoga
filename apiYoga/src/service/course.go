package service

import "api/db"

// 尝试service的另一种写法
// 新增一个课程
func (m *Message) InsertNewCourse(adminId, recommendMaxNum, recommendMinNum int, courseName, courseSubject, introduction, introductionURL string, isGroup, isTeam, isVIP bool) {
	err, isUnique := db.InsertNewCourse(adminId, recommendMaxNum, recommendMinNum, courseName, courseSubject, introduction, introductionURL, isGroup, isTeam, isVIP)
	if !isUnique {
		// 名字重复
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = false
	}
	if err != nil {
		// 其他错误
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
	}
	m.HaveError = false
	m.Info = "成功新增一个课程"
	m.IsSuccess = true
}

// 删除一个课程
func (m *Message) DropCourseByName(name string) {
	err, notExist := db.DeleteCourseByName(name)
	if notExist {
		// 课程不存在
		m.IsSuccess = false
		m.Info = "课程不存在"
		m.HaveError = false
	}
	if err != nil {
		// 其他错误
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
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
	}
	m.Result = courses
	m.HaveError = false
	m.Info = "获取课程列表成功"
	m.IsSuccess = true
}

//本文件中这些对于message的操作方法很好,能批量复制
