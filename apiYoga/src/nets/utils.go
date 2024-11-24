package nets

import (
	"api/service"

	"github.com/gin-gonic/gin"
)

// this is a util, if you handle the return of m is c.json the m.info you can do this
func returnMdotInfo(m service.Message, c *gin.Context) {
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	if !m.IsSuccess {
		c.JSON(400, gin.H{"message": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Info})
}

// similar to returnMdotInfo,but if success it will return m.result
func returnMdotInfoOrResult(m service.Message, c *gin.Context) {
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	if !m.IsSuccess {
		c.JSON(400, gin.H{"message": m.Info})
		return
	}
	c.JSON(200, m.Result)
}
