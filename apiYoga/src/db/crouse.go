package db

import (
	"errors"

	"api/config"
	"api/loger"
)

// 根据课程名称删除一个课程
func DeleteCourseByName(courseName string) (err error, isExist bool) {
	// 检查课程名称是否存在
	isExist = checkCourseNameExist(courseName)
	// 如果课程不存在，直接返回
	if !isExist {
		return nil, isExist
	}
	// 执行删除操作
	err = postdb.Where("course_name =?", courseName).Delete(&Course{}).Error
	if err != nil {
		// 记录错误日志
		loger.Loger.Println("error: DeleteCourseByName", err, "courseName:", courseName)
		return err, isExist
	}

	// 删除成功，返回 nil 错误和 notExist 标志
	return nil, isExist
}

// 新增一个课程,成功返回nil和true。失败返回error，可能返回.false名称重复
func InsertNewCourse(adminId, recommendMaxNum, recommendMinNum int, courseName, courseSubject, introduction, introductionURL string, isGroup, isTeam, isVIP bool) (err error, isUnique bool) {
	course := &Course{
		AdminId:         0,
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
	if !checkCourseNameExist(courseName) {
		return errors.New("课程名称重复,请另外挑选新的"), false
	}
	course.AdminId = adminId
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

// 检查课程是否名称是否存在
func checkCourseNameExist(courseName string) (isUnique bool) {
	var count int64
	err := postdb.Model(&Course{}).Where("course_name=?", courseName).Count(&count).Error
	if err != nil {
		loger.Loger.Println("error: 出现错误,找到课程名称为 ", courseName, "的课程时出现了错误", err.Error())
	}
	return count > 0
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
func SelectCourseInfo(courseId int) (course Course, err error) {
	err = postdb.Where("id=?", courseId).First(&course).Error
	if err != nil {
		loger.Loger.Println("error: 出现错误, 查询课程id为 ", courseId, "的课程时出现了错误", err.Error())
		return Course{}, err
	}
	return course, nil
}
func SelectCourseTypeByCourseId(courseId int) (courseType int, err error) {
	var course Course
	err = postdb.Where("id=?", courseId).First(&course).Error
	if err != nil {
		loger.Loger.Println("error: 出现错误, 查询课程id为 ", courseId, "的课程时出现了错误", err.Error())
		return 0, err
	}
	if course.IsVipType {
		return 0, nil
	}
	if course.IsTeamType {
		return 1, nil
	}
	if course.IsGroupType {
		return 2, nil
	}
	return -1, errors.New("课程错误，请联系管理员")
}
