package nets

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"api/db"
	"api/loger"
	"api/service"
)

// 第一个访问使用到的函数,token是不会更新的，所以返回的只有sessionId
func userLoginWithSessionAndToken(c *gin.Context) {
	type authenticationInfo struct {
		Session string `json:"session"`
		Token   string `json:"token"`
	}
	var getData authenticationInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(AuthenticationInfo{
		Session: getData.Session,
		Token:   getData.Token,
	}, c)
	if err != nil {
		return
	}
	var m service.Message
	m.SelectUserInfoByUserId(sessionInfo.UserId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	// 这个函数不存在没有错误还不成功的情况
	userInfo := m.Result.(db.User)
	c.JSON(200, gin.H{"session_id": sessionInfo.SessionId, "user_info": userInfo})
}
func userLoginWithCode(c *gin.Context) {
	type codeStruct struct {
		Code string `json:"code"`
	}
	var getData codeStruct
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var m service.Message
	m.RegisterUser(getData.Code)
	if m.IsSuccess {
		authenticationInfo, ok := m.Result.(db.Authentication)
		if ok {
			c.JSON(200, gin.H{"authentication": authenticationInfo})
			return
		} else {
			loger.Loger.Println("error:", "将用户信息转化的时候出现错误")
			c.JSON(400, gin.H{"error": "注册失败,请联系管理员 code:1"})
			return
		}
	} else {
		if m.HaveError {
			loger.Loger.Println("error:", m.Info)
			c.JSON(400, gin.H{"error": m.Info})
			return
		}
	}
}

// 注意，名字name和nickname不一样，name扮演者比较重要的角色，
func userRename(c *gin.Context) {
	type userInfo struct {
		NewName        string             `json:"newName"`
		Authentication AuthenticationInfo `json:"authentication"`
	}
	var getData userInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.Authentication, c)
	if err != nil {
		return
	}
	message := service.Rename(fmt.Sprint(sessionInfo.UserId), getData.NewName)
	if message.HaveError {
		c.JSON(400, gin.H{"error": message.Info})
	}
	c.JSON(200, gin.H{"message": "success"})
}

// 管理员和老师在这里登录，因为在gui所以不考虑时间损耗，使用token进行验证
func adminAndTeacherLogin(c *gin.Context) {
	type adminInfo struct {
		Level    int    `json:"level"`
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	var getData adminInfo

	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if getData.Level != 1 && getData.Level != 2 {
		c.JSON(400, gin.H{"message": "请输入合适的level,警告,不要进行抓包攻击,已经记录你的ip"})
	}
	var m service.Message
	m.AdminAndTeacherLogin(getData.Account, getData.Password, getData.Level)
	if m.HaveError {
		c.JSON(400, gin.H{"message": m.Info})
		return
	} else {
		if m.IsSuccess {
			token := m.Result.(string)
			c.JSON(200, gin.H{"token": token})
			return
		} else {
			c.JSON(400, gin.H{"message": m.Info})
			return
		}
	}
}

// electron 端的入口函数，每次进入都先调用一次
func electronEntrance(c *gin.Context) {
	type userInfo struct {
		Token string `json:"token"`
	}
	var getData userInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ok, htTeacher, account := authenticationInElectron(getData.Token, c)
	if !ok {
		return
	}
	var m service.Message
	m.SelectAdminOrTeacherInfoByAccount(account, htTeacher)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	if m.IsSuccess {
		c.JSON(200, m.Result)
	} else {
		c.JSON(400, gin.H{"message": m.Info})
	}
}

// 给微信端的用户更改
func updateUserInfo(c *gin.Context) {
	type newUserInfo struct {
		AuthenticationInfo AuthenticationInfo `json:"authentication"`
		Nickname           string             `json:"nickname"`
		Gender             bool               `json:"gender"`
		Signature          string             `json:"signature"`
	}
	var getData newUserInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	var m service.Message
	m.UpdateUserInfo(fmt.Sprint(sessionInfo.UserId), getData.Nickname, getData.Signature, getData.Gender)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
	}
	c.JSON(200, gin.H{"message": m.Info})
}

// 给教师修改自己的设置,目前只有这几个
func updateTeacherInfo(c *gin.Context) {
	type newTeacherInfo struct {
		Introduction       string             `json:"introduction"`
		AuthenticationInfo AuthenticationInfo `json:"authentication"`
	}
	var getData newTeacherInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sessionInfo, err := authentication(getData.AuthenticationInfo, c)
	if err != nil {
		return
	}
	var m service.Message
	m.UpdateTeacherInfo(fmt.Sprint(sessionInfo.UserId), getData.Introduction)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
	}
	c.JSON(200, gin.H{"message": m.Info})
}
