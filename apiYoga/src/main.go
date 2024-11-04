package main

import (
	"fmt"

	"api/config"
	"api/db"
	"api/loger"
	"api/nets"
	"api/service"
)

func main() {
	loger.StartApiYoga()
	config.UnmarshalConfig()
	db.StartClient()
	service.StartService()
	nets.StartApiEngine()
	update := make(chan struct{})
	//这里用来以后设置热更新功能
	<-update
	fmt.Println("Received stop signal")
}
