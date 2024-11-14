package nets

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"api/db"
	"api/loger"
	"api/service"
)

// electron 中管理员使用
func insertNewCourse(c *gin.Context) {
	type CourseInfo struct {
		RecommendMaxNum string `json:"recommendMaxNum"`
		RecommendMinNum string `json:"recommendMinNum"`
		CourseName      string `json:"courseName"`
		CourseSubject   string `json:"courseSubject"`
		Introduction    string `json:"introduction"`
		IntroductionURL string `json:"introductionURL"`
		IsGroup         bool   `json:"isIsGroup"`
		IsTeam          bool   `json:"isTeam"`
		IsVIP           bool   `json:"isVIP"`
		Token           string `json:"token"`
	}
	var getData CourseInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ok, htTeacher, account := authenticationInElectron(getData.Token, c)
	if !ok {
		return
	}
	if htTeacher == false {
		c.JSON(400, gin.H{"message": "只有管理员才能创建课程"})
		return
	}
	var m service.Message
	var recommendMaxNum int
	var recommendMinNum int
	recommendMaxNum, err := strconv.Atoi(getData.RecommendMaxNum)
	if err != nil {
		loger.Loger.Println("recommendMaxNum convert to int error: ", err)
		c.JSON(400, gin.H{"error": "recommendMaxNum 转换为 int 出错"})
		return
	}
	recommendMinNum, err = strconv.Atoi(getData.RecommendMinNum)
	if err != nil {
		loger.Loger.Println("recommendMinNum convert to int error: ", err)
		c.JSON(400, gin.H{"error": "recommendMinNum 转换为 int 出错"})
		return
	}
	m.InsertNewCourse(account, recommendMaxNum, recommendMinNum, getData.CourseName, getData.CourseSubject, getData.Introduction, getData.IntroductionURL, getData.IsGroup, getData.IsTeam, getData.IsVIP)
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
func deleteCourseByName(c *gin.Context) {
	type CourseInfo struct {
		Token      string `json:"token"`
		CourseName string `json:"course_name"`
	}
	var getData CourseInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, htTeacher, account := authenticationInElectron(getData.Token, c)
	if !isOk {
		return
	}
	if htTeacher == false {
		c.JSON(400, gin.H{"message": "只有管理员才能删除课程"})
		return
	}

	var m service.Message
	m.DropCourseByName(getData.CourseName)
	fmt.Println(m, "m2:")
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		fmt.Println("code3")
		return
	}
	if !m.IsSuccess {
		fmt.Println("不存在吧")
		fmt.Println("code2")
		c.JSON(400, gin.H{"message": m.Info})
		return
	}
	fmt.Println("code1")
	loger.Loger.Println("! dangerous ->", account, "<- 这个account的admin用户删除了一个课程:", getData.CourseName)
	c.JSON(200, gin.H{"message": m.Info})
}

// 检索课程列表,文本的传输并不需要什么资源消耗,完全可以作为一个大杂烩,想要任何课程内容都可以全部检索
func selectCourse(c *gin.Context) {
	type CourseInfo struct {
		AuthenticationInfo AuthenticationInfo `json:"authenticationInfo"`
	}
	var getData CourseInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := authentication(getData.AuthenticationInfo, c)
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
	courses, ok := m.Result.([]db.Course)
	if ok {
		c.JSON(200, gin.H{"courses": courses})
	}
}

// 事实上和上面这个函数的作用相同，内容也基本一致，但是显然两种验证方式的平台都有这个需求
func selectCourseByElectron(c *gin.Context) {
	type CourseInfo struct {
		Token string `json:"token"`
	}
	var getData CourseInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, _, _ := authenticationInElectron(getData.Token, c)
	if !isOk {
		return
	}
	var m service.Message
	m.SelectCourse()
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	var courses []db.Course
	courses, ok := m.Result.([]db.Course)
	if ok {
		c.JSON(200, gin.H{"courses": courses})
	}
}
