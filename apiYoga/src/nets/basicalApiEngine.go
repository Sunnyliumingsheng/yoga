package nets

import "github.com/gin-gonic/gin"

func basicalApiEngine(r *gin.Engine) {
	r.POST("/api/login", userLoginWithSessionAndToken)
	r.POST("/api/register", userLoginWithCode)
}
