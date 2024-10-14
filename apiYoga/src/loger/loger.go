package loger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var Loger *log.Logger

func init() {
	file := "../output/" + time.Now().Format("20180102") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Loger = log.New(logFile, "[qSkiptool]", log.LstdFlags|log.Lshortfile|log.LUTC) // 将文件设置为loger作为输出

}
func StartApiYoga() {
	fmt.Println("hello world")
	Loger.Println(time.Now(), "start ...")
}
