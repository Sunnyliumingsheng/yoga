package service

import (
	"api/db"
	"api/deamon"
	"time"
)

// 注意这里仅仅是添加了一个班级，并没有激活班级
func (m *Message) InsertNewClass(classList db.ClassList) {
	if classList.DayOfWeek < 0 || classList.DayOfWeek >= 7 {
		m.IsSuccess = false
		m.Info = "星期值不在0-6之间"
		m.HaveError = true
		return
	}
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

// select in the postgres not in the ram
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
	courseId, err := db.SelectCourseIdByClassId(classId)
	if err != nil {
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	isOK, err := CanStudentReserveThisCourse(userId, courseId)
	if err != nil {
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	if !isOK {
		m.IsSuccess = false
		m.Info = "该学生没有有效的卡，无法预约课程"
		m.HaveError = false
		return
	}
	err = deamon.Resume(userId, classId)
	if err != nil {
		m.IsSuccess = false
		m.Info = err.Error()
		m.HaveError = true
		return
	}
	m.IsSuccess = true
	m.Info = "预约成功"
	m.HaveError = false
}

// select in the ram
func (m *Message) SelectActivedClass() {
	fourDayClass, err := deamon.QuicklySelectClass()
	if err != nil {
		m.HaveError = true
		m.IsSuccess = false
		m.Info = err.Error()
		return
	}
	m.HaveError = false
	m.IsSuccess = true
	m.Result = fourDayClass
}
func (m *Message) SelectClassByClassId(classId int) {
	var classInfo db.ClassActived
	var resumeInfo []db.UserResumeInfo
	var err error
	classInfo, resumeInfo, err = deamon.QuicklySelectClassByClassId(classId)
	if err != nil {
		m.HaveError = true
		m.IsSuccess = false
		m.Info = err.Error()
		return
	}
	m.HaveError = false
	m.IsSuccess = true
	m.Result = map[string]interface{}{
		"classInfo":  classInfo,
		"resumeInfo": resumeInfo,
	}
}

// select in the postgres not the ram
func (m *Message) SelectTeachingClass(userId int) {
	teacherId, err := db.SelectTeacherIdByUserId(userId)
	if err != nil {
		m.HaveError = true
		m.IsSuccess = false
		m.Info = err.Error()
		return
	}
	weekday := int(time.Now().Weekday())
	class, err := db.SelectTeachClassThisWeekday(teacherId, weekday)
	if err != nil {
		m.HaveError = true
		m.IsSuccess = false
		m.Info = err.Error()
		return
	}
	m.HaveError = false
	m.IsSuccess = true
	m.Result = class
}
func (m *Message) SelectClass(classId int) {

}
