package nets

import (
	"errors"

	"github.com/gin-gonic/gin"

	"api/service"
	"api/util"
)

type AuthenticationInfo struct {
	Session string `json:"session"`
	Token   string `json:"token"`
}

// 用于日常的身份验证,这是非常常用的函数,每个函数都会用到,如果有错误只需要直接return就行了,不需要自己回复前端
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

// 这里本来有一个验证admin的函数的，现在不需要了，我认为以admin身份进行操作的地方应该在electron写的gui里面
// 上课学生签到和老师核验应该在微信小程序里进行，老师的高级操作和管理员的日常操作应该在elecron的gui里面进行
// 也就是说
// 这个用在小程序里面进行验证拥有老师以上的权限
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

// 这个用在relecton中用来验证拥有老师以上的权限higher than teacher,使用的时候，如果不ok就return，ok就接着看
func authenticationInElectron(token string, c *gin.Context) (ok bool, htTeacher bool, account string) {
	account, level, err := util.ParseTokenWithLevelAndAccount(token)
	if err != nil {
		c.JSON(400, gin.H{"error": "验证失败，请重新登录"})
		return false, false, ""
	}
	if level == 2 {
		return true, false, account
	}
	return true, true, account
}
