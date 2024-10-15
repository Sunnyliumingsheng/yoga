package main

import (
	"cli/cmd"
	"cli/config"
)

func main() {
	config.UnmarshalConfig()
	cmd.Execute()
}
