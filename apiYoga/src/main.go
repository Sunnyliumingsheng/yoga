package main

import (
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
}
