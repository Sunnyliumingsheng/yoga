package nets

import "github.com/gin-gonic/gin"

func basicalApiEngine(r *gin.Engine) {
	// 用户端使用的  有关用户的 api集合
	r.POST("/api/login", userLoginWithSessionAndToken)
	r.POST("/api/register", userLoginWithCode)
	r.POST("/api/rename", userRename)
	// 开发者和管理员使用的  有关用户的 api集合
	r.POST("/api/root/register/admin", sudoRegisterAdmin)
	r.POST("/api/root/login", sudoLogin)
	r.POST("/api/root/select/user/by/name", selectUserInfoByName)
	r.POST("/api/root/insert/user/by/name", insertNewUser)
	r.POST("/api/root/drop/user/by/userId", dropUserByUserId)
	r.POST("/api/root/update/user/level/by/name", updateUserLevel)
	// 公共api集合
}
