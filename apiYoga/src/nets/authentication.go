package nets

import (
	"errors"

	"github.com/gin-gonic/gin"

	"api/service"
)

type AuthenticationInfo struct {
	Session string `json:"session"`
	Token   string `json:"token"`
}

// 用于日常的身份验证,这是非常常用的函数,每个函数都会用到
func authentication(info AuthenticationInfo, c *gin.Context) (err error) {
	message := service.SessionAndTokenAuthentication(info.Session, info.Token)
	if message.HaveError {
		c.JSON(400, gin.H{"error": "比较严重的登录错误,请截图并联系管理员"})
		return errors.New(message.Info)
	}
	if !message.IsSuccess {
		c.JSON(200, gin.H{"message": "重新登录"})
		return errors.New(message.Info)
	}
	return nil
}
func authenticationAdmin(info AuthenticationInfo, c *gin.Context) (ok bool) {
	message := service.SessionAndTokenAuthentication(info.Session, info.Token)
	if message.HaveError {
		c.JSON(400, gin.H{"error": "比较严重的登录错误,请截图并联系管理员"})
		return false
	}
	if !message.IsSuccess {
		c.JSON(200, gin.H{"message": "重新登录"})
		return false
	}
	level, _ := message.Result.(int)
	if level <= 1 {
		return true
	}
	return false

}
func authenticationTeacher(info AuthenticationInfo, c *gin.Context) (ok bool) {
	message := service.SessionAndTokenAuthentication(info.Session, info.Token)
	if message.HaveError {
		c.JSON(400, gin.H{"error": "比较严重的登录错误,请截图并联系管理员"})
		return false
	}
	if !message.IsSuccess {
		c.JSON(200, gin.H{"message": "重新登录"})
		return false
	}
	level, _ := message.Result.(int)
	if level <= 2 {
		return true
	}
	return false

}
