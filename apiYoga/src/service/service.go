package service

import (
	"time"

	"api/loger"
)

func StartService() {
	loger.Loger.Println(time.Now(), "启动第二层后端逻辑处理服务")
}
