package nets

import (
	"github.com/gin-gonic/gin"

	"api/config"
	"api/service"
)

type SudoAuthentication struct {
	SudoName     string `json:"account"`
	SudoPassword string `json:"password"`
}

// 如果你想修改登录逻辑,请修改这里,修改之前,请确定你已经看了config/types.go
func authenticateSudo(s SudoAuthentication) bool {
	return s.SudoName == config.Config.Sudo.SuperUsername && s.SudoPassword == config.Config.Sudo.SuperPassword
}

func sudoLogin(c *gin.Context) {
	type loginInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
	}
	var getData loginInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if authenticateSudo(getData.SudoAuthentication) {
		c.JSON(200, gin.H{"message": "success"})
	} else {
		c.JSON(400, gin.H{"message": "wrong account or password"})
	}
}
func sudoRegisterAdmin(c *gin.Context) {
	type registerAdminInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		UserName           string             `json:"userName"`
		AdminAccount       string             `json:"adminAccount"`
		AdminPassword      string             `json:"adminPassword"`
	}
	var getData registerAdminInfo
	if authenticateSudo(getData.SudoAuthentication) {
		c.JSON(200, gin.H{"message": "success"})
	} else {
		c.JSON(400, gin.H{"message": "wrong account or password"})
	}
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

}

// 通过用户姓名对其信息进行检索
func selectUserInfoByName(c *gin.Context) {
	type UserInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		Name               string             `json:"name"`
	}
	var getData UserInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	message := service.SelectUserInfoByName(getData.Name)
	if message.HaveError {
		c.JSON(200, gin.H{"error": "查询出现错误"})
		return
	}
	if message.IsSuccess {
		c.JSON(200, message.Result)
		return
	} else {
		c.JSON(200, gin.H{"message": "没有这个用户,请确认是用户名而不是昵称"})
		return
	}
}

// 仅供测试使用,所以openid是空的,只有name就可以创建
func insertNewUser(c *gin.Context) {
	type UserInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		Name               string             `json:"name"`
	}
	var getData UserInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	message := service.InsertNewUser(getData.Name)
	if message.HaveError {
		c.JSON(200, gin.H{"error": message.Info})
		return
	}
	c.JSON(200, gin.H{"message": message.Info})
}
func dropUserByUserId(c *gin.Context) {
	type UserInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		UserId             string             `json:"userId"`
	}
	var getData UserInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	message := service.DropUserByStringUserId(getData.UserId)
	if message.HaveError {
		c.JSON(200, gin.H{"error": message.Info})
		return
	}
	c.JSON(200, gin.H{"message": message.Info})
}
