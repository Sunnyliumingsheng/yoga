package nets

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"api/loger"
)

var r *gin.Engine

func init() {
	r = gin.New()
	logfile := "../output/" + time.Now().String() + "gin.log"
	f, err := os.Create(logfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 将日志输出位置设置为输出到文件
	r.Use(gin.LoggerWithWriter(f))
	r.Use()
	r.GET("/ping1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go addtionalApi()
	r.Run()
}
func StartApiEngine() {
	loger.Loger.Println(time.Now(), "start api engine server")
}
func addtionalApi() {
	time.Sleep(3 * time.Second)
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
