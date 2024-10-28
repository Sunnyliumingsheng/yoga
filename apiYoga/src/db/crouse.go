package db

import (
	"errors"

	"api/config"
	"api/loger"
)

// 新增一个课程,成功返回nil和true。失败返回error，可能返回.false名称重复
func InsertNewCourse(adminId, recommendMaxNum, recommendMinNum int, courseName, courseSubject, introduction, introductionURL string, isGroup, isTeam, isVIP bool) (err error, isUnique bool) {
	course := &Course{
		AdminID:         0,
		CourseName:      config.Config.NewCourseDefaultInfo.CourseName,
		CourseSubject:   config.Config.NewCourseDefaultInfo.CourseSubject,
		IntroductionURL: config.Config.NewCourseDefaultInfo.Introduction,
		Introduction:    config.Config.NewCourseDefaultInfo.IntroductionURL,
		IsGroupType:     true,
		IsTeamType:      false,
		IsVIPType:       false,
		RecommendMaxNum: recommendMaxNum,
		RecommendMinNum: recommendMinNum,
	}
	if !checkCourseNameUnique(courseName) {
		return errors.New("课程名称重复,请另外挑选新的"), false
	}
	course.AdminID = adminId
	course.CourseName = courseName
	course.CourseSubject = courseSubject
	course.Introduction = introduction
	course.IntroductionURL = introductionURL
	course.IsGroupType = isGroup
	course.IsTeamType = isTeam
	course.IsVIPType = isVIP
	course.RecommendMaxNum = recommendMaxNum
	course.RecommendMinNum = recommendMinNum

	err = postdb.Create(course).Error
	if err != nil {
		loger.Loger.Println("error: 出现错误, 创建课程时出现了错误", err.Error())
		return err, false
	}
	return nil, true
}
func checkCourseNameUnique(courseName string) (isUnique bool) {
	var count int64
	err := postdb.Model(&Course{}).Where("course_name=?", courseName).Count(&count).Error
	if err != nil {
		loger.Loger.Println("error: 出现错误,找到课程名称为 ", courseName, "的课程时出现了错误", err.Error())
	}
	return count == 0
}
