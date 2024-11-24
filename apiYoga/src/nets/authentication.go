package nets

import (
	"errors"

	"github.com/gin-gonic/gin"

	"api/config"
	"api/service"
	"api/util"
)

type AuthenticationInfo struct {
	Session string `json:"session"`
	Token   string `json:"token"`
}

// 唉，解藕的好处啊
// 如果你想修改超级用户的登录逻辑,请修改这里,修改之前,请确定你已经看了config/types.go
func authenticateSudo(s SudoAuthentication) bool {
	return s.SudoName == config.Config.Sudo.SuperUsername && s.SudoPassword == config.Config.Sudo.SuperPassword
}

// 用于日常的身份验证,这是非常常用的函数,每个函数都会用到,如果有错误只需要直接return就行了,不需要自己回复前端
func authentication(info AuthenticationInfo, c *gin.Context) (sessionInfo service.SessionInfo, err error) {
	var m service.Message
	m.SessionAndTokenAuthentication(info.Session, info.Token)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return sessionInfo, errors.New("存在错误" + m.Info)
	}
	if !m.IsSuccess {
		c.JSON(400, gin.H{"message": m.Info})
		return sessionInfo, errors.New("验证失败" + m.Info)
	}
	sessionInfo = m.Result.(service.SessionInfo)
	return sessionInfo, nil
}

// 这里本来有一个验证admin的函数的，现在不需要了，我认为以admin身份进行操作的地方应该在electron写的gui里面
// 上课学生签到和老师核验应该在微信小程序里进行，老师的高级操作和管理员的日常操作应该在elecron的gui里面进行
// 也就是说
// 这个用在小程序里面进行验证拥有老师以上的权限
func authenticationTeacher(info AuthenticationInfo, c *gin.Context) (ok bool) {
	var m service.Message
	m.SessionAndTokenAuthentication(info.Session, info.Token)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	if !m.IsSuccess {
		c.JSON(400, gin.H{"message": m.Info})
		return
	}
	sessionInfo := m.Result.(service.SessionInfo)
	level := sessionInfo.Level
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
