package nets

import (
	"api/db"
	"api/deamon"
	"api/service"

	"github.com/gin-gonic/gin"
)

func insertNewClass(c *gin.Context) {
	type classInfo struct {
		ClassName     string `json:"class_name"`
		CourseId      int    `json:"course_id"`
		Auto          bool   `json:"auto"`
		DayOfWeek     int    `json:"day_of_week"`
		AlreadyActive bool   `json:"already_active"`
		Index         int    `json:"index"`
		Min           int    `json:"min"`
		Max           int    `json:"max"`
		TeacherId     int    `json:"teacher_id"`
		Token         string `json:"token"`
	}
	var getData classInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, isAdmin, _ := authenticationInElectron(getData.Token, c)
	if !isOk {
		return
	}
	if !isAdmin {
		c.JSON(400, gin.H{"message": "只有管理员才能创建新班级"})
		return
	}
	var classList db.ClassList
	classList.ClassName = getData.ClassName
	classList.CourseId = getData.CourseId
	classList.Auto = getData.Auto
	classList.DayOfWeek = getData.DayOfWeek
	classList.AlreadyActive = getData.AlreadyActive
	classList.Index = getData.Index
	classList.Min = getData.Min
	classList.Max = getData.Max
	classList.TeacherId = getData.TeacherId
	var m service.Message
	m.InsertNewClass(classList)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Info})
}
func deleteClass(c *gin.Context) {
	type deleteInfo struct {
		ClassId int    `json:"class_id"`
		Token   string `json:"token"`
	}
	var getData deleteInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, isAdmin, _ := authenticationInElectron(getData.Token, c)
	if !isOk {
		return
	}
	if !isAdmin {
		c.JSON(400, gin.H{"message": "只有管理员才能删除课程"})
		return
	}
	var m service.Message
	m.DeleteClass(getData.ClassId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Info})
}

// electron can active class
func activeClass(c *gin.Context) {
	type classInfo struct {
		ClassId int    `json:"class_id"`
		Token   string `json:"token"`
	}
	var getData classInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, isAdmin, _ := authenticationInElectron(getData.Token, c)
	if !isOk {
		return
	}
	if !isAdmin {
		c.JSON(400, gin.H{"message": "只有管理员才能更改课程状态"})
		return
	}
	var m service.Message
	m.ActiveClass(getData.ClassId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Info})
}

// electron can do this
func selectAllClass(c *gin.Context) {
	type selectInfo struct {
		Token string `json:"token"`
	}
	var getData selectInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, _, _ := authenticationInElectron(getData.Token, c)
	if !isOk {
		return
	}
	var m service.Message
	m.SelectAllClass()
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	var classList []db.ClassList
	classList, ok := m.Result.([]db.ClassList)
	if !ok {
		c.JSON(400, gin.H{"error": "解析返回结果失败"})
		return
	}
	c.JSON(200, gin.H{"message": classList})
}

// 任何人都可以使用 in the postgres
func selectAllActivedClass(c *gin.Context) {
	var m service.Message
	m.SelectAllActivedClass()
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	var classList []db.ClassList
	classList, ok := m.Result.([]db.ClassList)
	if !ok {
		c.JSON(400, gin.H{"error": "解析返回结果失败"})
		return
	}
	c.JSON(200, gin.H{"message": classList})
}

// select in the ram,everyone can do this
func selectActivedClass(c *gin.Context) {
	var m service.Message
	m.SelectActivedClass()
	c.JSON(200, m.Result)
}

// select in ram,everyone can do this
func selectClassByClassId(c *gin.Context) {
	type SelectInfo struct {
		ClassId int `json:"class_id"`
	}
	var getData SelectInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var m service.Message
	m.SelectClassByClassId(getData.ClassId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, m.Result)
}

func resume(c *gin.Context) {
	type userInfo struct {
		AuthenticationInfo AuthenticationInfo `json:"authentication"`
		ClassId            int                `json:"class_id"`
	}
	var getData userInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	if sessionInfo.Level == 4 || sessionInfo.Level > 3 {
		c.JSON(400, gin.H{"message": "请先寻找管理员注册为正式学员"})
	}
	var m service.Message
	m.Resume(getData.ClassId, sessionInfo.UserId)
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

// select in postgres ,teacher use in wx
func selectTeachingClass(c *gin.Context) {
	type selectInfo struct {
		AuthenticationInfo AuthenticationInfo `json:"authentication"`
	}
	var getData selectInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	if sessionInfo.Level > 3 {
		c.JSON(400, gin.H{"message": "请先成为教师"})
	}
	userId := sessionInfo.UserId
	var m service.Message
	m.SelectTeachingClass(userId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	if !m.IsSuccess {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, m.Result)
}

// student can do this
func cancelResume(c *gin.Context) {
	type cancelInfo struct {
		AuthenticationInfo AuthenticationInfo `json:"authenticationInfo"`
		ClassId            int                `json:"classId"`
	}
	var getData cancelInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	var m service.Message
	m.CancelResume(sessionInfo.UserId, getData.ClassId)
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

// student can do this
func selectMyResume(c *gin.Context) {
	type userInfo struct {
		AuthenticationInfo AuthenticationInfo `json:"authentication"`
	}
	var getData userInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	if sessionInfo.Level > 3 {
		c.JSON(400, gin.H{"error": "请先申请成为学生"})
		return
	}
	var m service.Message
	m.SelectMyResume(sessionInfo.UserId)
	c.JSON(200, m.Result)
}

// change a user checkin status , teacher can do this
func updateTeacherInfoCheckinStatusByUserId(c *gin.Context) {
	type checkinInfo struct {
		UserId             int                `json:"user_id"`
		AuthenticationInfo AuthenticationInfo `json:"authenticationInfo"`
		CheckinStatus      int                `json:"checkin_status"`
		ClassId            int                `json:"class_id"`
	}
	var getData checkinInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	if sessionInfo.Level > 2 {
		c.JSON(400, gin.H{"error": "没有权限"})
		return
	}
	var m service.Message
	m.ChangeCheckinStatusUser(getData.UserId, getData.CheckinStatus, getData.ClassId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
	}
	c.JSON(200, gin.H{"message": "success"})
}

// checkin all user , cause it is usually to use
func checkinAllStudent(c *gin.Context) {
	type checkinInfo struct {
		ClassId            int                `json:"class_id"`
		AuthenticationInfo AuthenticationInfo `json:"authenticationInfo"`
		Text               string             `json:"text"`
	}
	var getData checkinInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	if sessionInfo.Level > 2 {
		c.JSON(400, gin.H{"error": "没有权限"})
	}
	var m service.Message
	m.CheckinAllStudent(sessionInfo.UserId, getData.ClassId, getData.Text)
	returnMdotInfo(m, c)
}

// select all record ,everyone can do this
func selectRecord(c *gin.Context) {
	type recordLength struct {
		Tail int `json:"tail"`
	}
	var getData recordLength
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var m service.Message
	m.SelectRecord(getData.Tail)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, m.Result)
}

// cli can do this ,select the black list
func selectBlackListByUserId(c *gin.Context) {
	type selectInfo struct {
		UserId             int                `json:"user_id"`
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
	}
	var getData selectInfo
	isOK := authenticateSudo(getData.SudoAuthentication)
	if !isOK {
		c.JSON(400, gin.H{"error": "验证失败"})
		return
	}
	yes, err := db.IsSomeoneInBlackList(getData.UserId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"is_in_black_list": yes})
}

// 附赠的说明，这里很简陋
// remove someone out of blask list ,only sudo can do this
func deleteAllBlackList(c *gin.Context) {
	type authInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
	}
	var getData authInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOK := authenticateSudo(getData.SudoAuthentication)
	if !isOK {
		c.JSON(400, gin.H{"error": "验证失败"})
		return
	}
	err := db.DeleteAllBlackList()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
func SelectBlackList(c *gin.Context) {
	type authInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
	}
	var getData authInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOK := authenticateSudo(getData.SudoAuthentication)
	if !isOK {
		c.JSON(400, gin.H{"error": "验证失败"})
		return
	}
	list, err := db.SelectBlackList()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"list": list})
}

// if you want to shutdown -r but still storage the ram,you can sudo this and quickly reboot
func storageRam(c *gin.Context) {
	deamon.Storage()
}
func getStorageRam(c *gin.Context) {
	deamon.GetStorage()
}
