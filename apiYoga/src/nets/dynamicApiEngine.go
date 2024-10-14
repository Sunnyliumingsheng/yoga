package nets

import (
	"time"

	"github.com/gin-gonic/gin"
)

func dynamicApi() {
	time.Sleep(3 * time.Second)
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
