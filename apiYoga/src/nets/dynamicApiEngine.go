package nets

import (
	"api/loger"
	"time"

	"github.com/gin-gonic/gin"
)

func dynamicApi() {
	time.Sleep(1 * time.Second)
	loger.Loger.Println("dynamic api running... ")
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
