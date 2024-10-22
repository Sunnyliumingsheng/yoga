package main

import (
	"cli/config"
	"cli/handler"
)

func main() {
	config.UnmarshalConfig()
	handler.HandleFlags()
}
