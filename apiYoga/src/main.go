package main

import (
	"log"
	"time"

	"api/config"
	"api/db"
	"api/service"
)

func main() {
	log.Println(time.Now(), "Start running new pro")
	config.UnmarshalConfig()
	db.StartClient()
	service.StartServer()
}
