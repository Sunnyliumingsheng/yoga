package nets

import (
	"github.com/gin-gonic/gin"

	"api/config"
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
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

}