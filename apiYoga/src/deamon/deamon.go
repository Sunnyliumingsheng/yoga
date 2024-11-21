package deamon

import (
	"api/loger"
	"time"
)

func StartAllDeamon() {
	loger.Loger.Println("finish all deamon")
	immediatelyToDo()
	go nightFlash()
}

// 每天凌晨3点刷新一次
func nightFlash() {
	now := time.Now()
	nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 3, 0, 0, 0, now.Location())
	duration := nextMidnight.Sub(now)
	time.Sleep(duration)
	nightToDo()
	for range time.Tick(24 * time.Hour) {
		nightToDo()
	}
}
func nightToDo() {
	loger.Loger.Println("新的一天刷新内存", time.Now())
	FlashUserCard()
	FlashActivedClass()
}
func immediatelyToDo() {
	FlashUserCard()
}
