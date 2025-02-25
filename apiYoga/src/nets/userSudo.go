package nets

import (
	"github.com/gin-gonic/gin"

	"api/service"
)

// 在这里的所有函数,都只需要返回m.result就行了,文本格式就是返回到命令汉的格式,不是json

type SudoAuthentication struct {
	SudoName     string `json:"account"`
	SudoPassword string `json:"password"`
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
func updateUserLevel(c *gin.Context) {
	type userInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		Name               string             `json:"name"`
		IsStudent          bool               `json:"isStudent"`
		IsTeacher          bool               `json:"isTeacher"`
		IsAdmin            bool               `json:"isAdmin"`
	}
	var getData userInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	message := service.UpdateUserLevel(getData.Name, getData.IsStudent, getData.IsTeacher, getData.IsAdmin)
	if message.HaveError {
		c.JSON(200, gin.H{"error": message.Info})
		return
	}
	c.JSON(200, gin.H{"message": message.Info})
}
func selectUserTail(c *gin.Context) {
	type TailInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		Tail               int                `json:"tail"`
	}
	var getData TailInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	var m service.Message
	m.SelectUserTail(getData.Tail)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	if !m.IsSuccess {
		c.JSON(400, gin.H{"message": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Result})
}

// 给者老师设置账号和密码
func insertTeacherAccountAndPassword(c *gin.Context) {
	type TeacherInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		Account            string             `json:"account"`
		Password           string             `json:"password"`
		TeacherId          int                `json:"teacherId"`
	}
	var getData TeacherInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	var m service.Message
	m.InsertTeacherAccountAndPassword(getData.TeacherId, getData.Account, getData.Password)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	} else {
		if m.IsSuccess {
			c.JSON(200, gin.H{"message": m.Info})
			return
		} else {
			c.JSON(200, gin.H{"message": m.Info})
			return
		}
	}
}

// 给管理员设置账号和密码
func insertAdminAccountAndPassword(c *gin.Context) {
	type AdminInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		Account            string             `json:"account"`
		Password           string             `json:"password"`
		AdminId            int                `json:"adminId"`
	}
	var getData AdminInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	var m service.Message
	m.InsertAdminAccountAndPassword(getData.AdminId, getData.Account, getData.Password)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	} else {
		if m.IsSuccess {
			c.JSON(200, gin.H{"message": m.Info})
			return
		} else {
			c.JSON(200, gin.H{"message": m.Info})
			return
		}
	}
}

// 查询管理员信息
func selectAdminInfoByName(c *gin.Context) {
	type AdminInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		Name               string             `json:"name"`
	}
	var getData AdminInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	var m service.Message
	m.SelectAdminInfoByName(getData.Name)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	} else {
		if m.IsSuccess {
			c.JSON(200, gin.H{"message": m.Result})
			return
		} else {
			//没有错误但是不成功就是没检索到
			c.JSON(200, gin.H{"message": m.Info})
			return
		}
	}
}

// 查询教师信息
func selectTeacherInfo(c *gin.Context) {
	type TeacherInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		Name               string             `json:"name"`
	}
	var getData TeacherInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"message": "wrong account or password"})
		return
	}
	var m service.Message
	m.SelectTeacherInfoByName(getData.Name)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	} else {
		if m.IsSuccess {
			c.JSON(200, gin.H{"message": m.Result})
			return
		} else {
			//没有错误但是不成功就是没检索到
			c.JSON(200, gin.H{"message": m.Info})
			return
		}
	}
}
