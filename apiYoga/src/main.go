package main

import (
	"api/config"
	"api/db"
	"api/deamon"
	"api/loger"
	"api/nets"
	"api/picture"
	"api/service"
)

func main() {
	picture.StartStoragePicture()
	loger.StartApiYoga()
	config.UnmarshalConfig()
	db.StartClient()
	service.StartService()
	deamon.StartAllDeamon()
	nets.StartApiEngine()
}
