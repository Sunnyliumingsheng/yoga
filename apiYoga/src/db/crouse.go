package db

import (
	"errors"

	"api/config"
	"api/loger"
)

// 根据课程名称删除一个课程
func DeleteCourseByName(courseName string) (err error, notExist bool) {
	course := &Course{}
	notExist = checkCourseNameUnique(courseName)
	if notExist {
		return nil, notExist
	}
	err = postdb.Where("course_name =?", courseName).Delete(course).Error
	if err != nil {
		loger.Loger.Println("error: DeleteCourseByName", err, "courseName:", courseName)
		return err, notExist
	}
	return nil, notExist
}

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
		IsVipType:       false,
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
	course.IsVipType = isVIP
	course.RecommendMaxNum = recommendMaxNum
	course.RecommendMinNum = recommendMinNum

	err = postdb.Model(&Course{}).Create(course).Error
	if err != nil {
		loger.Loger.Println("error: 出现错误, 创建课程时出现了错误", err.Error())
		return err, false
	}
	return nil, true
}

// 检查课程是否名称唯一
func checkCourseNameUnique(courseName string) (isUnique bool) {
	var count int64
	err := postdb.Model(&Course{}).Where("course_name=?", courseName).Count(&count).Error
	if err != nil {
		loger.Loger.Println("error: 出现错误,找到课程名称为 ", courseName, "的课程时出现了错误", err.Error())
	}
	return count == 0
}

// 检索所有课程
func SelectCourse() (courses []Course, err error) {
	err = postdb.Find(&courses).Error
	if err != nil {
		loger.Loger.Println("error: 出现错误, 查询所有课程时出现了错误", err.Error())
		return nil, err
	}
	return courses, nil
}
