package nets

import (
	"github.com/gin-gonic/gin"

	"api/db"
	"api/loger"
	"api/service"
)

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

	message := service.SessionAndTokenAuthentication(getData.Session, getData.Token)
	if message.IsSuccess {
		c.JSON(200, gin.H{"message": "success"})
	} else {
		if message.HaveError {
			result, ok := message.Result.(string)
			if ok {
				loger.Loger.Println("error:", message.Info, result)
				c.JSON(400, gin.H{"error": result})
			} else {
				loger.Loger.Println(message.Info)
				c.JSON(400, gin.H{"error": "出现错误,请联系管理员"})
			}
		} else {
			c.JSON(400, gin.H{"message": "session and token are expired both"})
		}
	}
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
	message := service.RegisterUser(getData.Code)
	if message.IsSuccess {
		authenticationInfo, ok := message.Result.(db.Authentication)
		if ok {
			c.JSON(200, gin.H{"message": authenticationInfo})
		} else {
			c.JSON(400, gin.H{"error": "注册失败,请联系管理员"})
		}
	} else {
		if message.HaveError {
			result, ok := message.Result.(string)
			if ok {
				loger.Loger.Println("error:", message.Info, result)
				c.JSON(400, gin.H{"error": result})
			} else {
				loger.Loger.Println(message.Info)
				c.JSON(400, gin.H{"error": "注册出现错误,请联系管理员"})
			}
		} else {
			c.JSON(400, gin.H{"message": "出现错误,请联系管理员"})
		}
	}
}

// 通过用户姓名对其信息进行检索
func selectUserInfoByName(c *gin.Context) {
	type UserInfo struct {
		Name string `json:"name"`
	}
	var getData UserInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	message := service.SelectUserInfoByName(getData.Name)
	if message.IsSuccess {
		c.JSON(200, gin.H{"userInfo": message.Result})
		return
	}
	if message.HaveError {

	}
}
