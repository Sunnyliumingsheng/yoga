package nets

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"api/db"
	"api/loger"
	"api/service"
)

func insertNewCourse(c *gin.Context) {
	type CourseInfo struct {
		RecommendMaxNum    int
		RecommendMinNum    int
		CourseName         string
		CourseSubject      string
		Introduction       string
		IntroductionURL    string
		IisGroup           bool
		IsTeam             bool
		IsVIP              bool
		AuthenticationInfo AuthenticationInfo
	}
	var getData CourseInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ok := authenticationAdmin(getData.AuthenticationInfo, c)
	if !ok {
		c.JSON(400, gin.H{"message": "验证失败"})
		return
	}
	var m service.Message
	userId, err := strconv.Atoi(getData.AuthenticationInfo.Session)
	if err != nil {
		loger.Loger.Println("error:", err, "session:", getData.AuthenticationInfo.Session)
		c.JSON(400, gin.H{"error": "错误,请联系开发者"})
		return
	}
	m.InsertNewCourse(userId, getData.RecommendMaxNum, getData.RecommendMinNum, getData.CourseName, getData.CourseSubject, getData.Introduction, getData.IntroductionURL, getData.IisGroup, getData.IsTeam, getData.IsVIP)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	if !m.IsSuccess {
		c.JSON(400, gin.H{"message": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Info})
}
func dropCourseByName(c *gin.Context) {
	type CourseInfo struct {
		AuthenticationInfo AuthenticationInfo
		CourseName         string
	}
	var getData CourseInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ok := authenticationAdmin(getData.AuthenticationInfo, c)
	if !ok {
		c.JSON(400, gin.H{"message": "验证失败权限不够"})
		return
	}
	var m service.Message
	m.DropCourseByName(getData.CourseName)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	if !m.IsSuccess {
		c.JSON(400, gin.H{"message": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Info})
}

// 检索课程列表,文本的传输并不需要什么资源消耗,完全可以作为一个大杂烩,想要任何课程内容都可以全部检索
func SelectCourse(c *gin.Context) {
	type CourseInfo struct {
		AuthenticationInfo AuthenticationInfo
	}
	var getData CourseInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	var m service.Message
	m.SelectCourse()
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	var courses []db.Course
	courses,ok:=m.Result.([]db.Course)
	if ok{
		c.JSON(200, gin.H{"courses": courses})
	}
}
