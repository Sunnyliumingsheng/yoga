package nets

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"api/loger"
)

var r *gin.Engine

// 入口函数, 最开始这里是init函数，但是发现使用init会有一些日志打印上的问题，现在使用这个就没有问题了
// 推测是因为init会在引入的时候加载，所以会造成堵塞，config，日志等配置都不会加载，就等着接受请求了
func entrance() {
	logfile := "../output/" + time.Now().Format("2006-01-02_15-04-05") + "_gin.log"
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 设置 gin 的日志输出到文件
	gin.DefaultWriter = io.MultiWriter(f)

	r = gin.New()
	r.Use(gin.LoggerWithWriter(f)) // 使用文件作为日志输出
	r.Use(gin.Recovery())          // 启用恢复中间件

	r.GET("/ping1", func(c *gin.Context) {
		loger.Loger.Println("ping1 api get request... ")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 注册其他路由
	basicalApiEngine(r)
	go dynamicApi() // 启动动态 API
	r.Run(":8080")  // 启动 Gin 服务器
}

func StartApiEngine() {
	loger.Loger.Println(time.Now(), "start api engine server")
	entrance()
}
