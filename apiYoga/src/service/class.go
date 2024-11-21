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
func (m *Message) ActiveClass(classId int) {
	err := db.ActiveClass(classId)
	if err != nil {
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	m.IsSuccess = true
	m.Info = "激活课程成功"
	m.HaveError = false
	return
}
func (m *Message) SelectAllClass() {
	classList, err := db.SelectAllClass()
	if err != nil {
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	m.IsSuccess = true
	m.Info = "查询所有课程成功"
	m.Result = classList
}
func (m *Message) SelectAllActivedClass() {
	classList, err := db.SelectAllActivedClass()
	if err != nil {
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	m.IsSuccess = true
	m.Info = "查询所有已激活课程成功"
	m.Result = classList
}
func (m *Message) Resume(classId, userId int) {
	// first check the user own a useful card

	// check this
}
