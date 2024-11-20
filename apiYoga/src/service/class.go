package service

import "api/db"

// 注意这里仅仅是添加了一个班级，并没有激活班级
func (m *Message) InsertNewClass(classList db.ClassList) {
	_, err := db.InsertNewClass(classList)
	if err != nil {
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}

	m.IsSuccess = true
	m.Info = "添加课程成功"
	m.HaveError = false
	return
}
func (m *Message) DeleteClass(classId int) {
	err := db.DeleteClass(classId)
	if err != nil {
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	m.IsSuccess = true
	m.Info = "删除课程成功"
	m.HaveError = false
	return
}
