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

	<-update
	fmt.Println("Received stop signal")
}
