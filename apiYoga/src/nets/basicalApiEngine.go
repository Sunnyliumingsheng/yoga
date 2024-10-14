package nets

import "github.com/gin-gonic/gin"

func basicalApiEngine(r *gin.Engine) {
	// 用户端使用的  有关用户的 api集合
	r.POST("/api/login", userLoginWithSessionAndToken)
	r.POST("/api/register", userLoginWithCode)
	// 开发者使用的  有关用户的 api集合
	r.POST("/api/root/register/admin", sudoRegisterAdmin)
}
