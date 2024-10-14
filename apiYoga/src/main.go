package main

import (
	"fmt"
	"time"

	"api/config"
	"api/db"
	"api/loger"
	"api/nets"
	"api/service"
)

func main() {
	loger.StartApiYoga()
	fmt.Println("hello!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	config.UnmarshalConfig()
	db.StartClient()
	service.StartService()
	time.Sleep(3 * time.Second)
	nets.StartApiEngine()
	update := make(chan struct{})

	<-update
	fmt.Println("Received stop signal")
}
