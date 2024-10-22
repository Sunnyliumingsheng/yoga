package main

import (
	"cli/config"
	"cli/handler"
	"cli/nets"
)

func main() {
	config.UnmarshalConfig()
	nets.Init()

	handler.HandleFlags()
}
