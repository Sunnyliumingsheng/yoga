package service

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.Default()
	go deamonManager(r)
	go startApi(r)

}

func deamonManager(r *gin.Engine) {
	r.GET("/api/test/"+"manage", testdeamon)
}
func startApi(r *gin.Engine) {
	r.GET("/api/ping", testapi)
	r.Run(":8080")
}
func StartServer() {
	log.Println("启动后端api服务", time.Now())
}
func testapi(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func testdeamon(c *gin.Context) {
	c.JSON(200, gin.H{
		"deamonName": "sucess",
	})
}
