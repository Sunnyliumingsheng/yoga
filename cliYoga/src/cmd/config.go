package cmd

import (
	"cli/config"
)

func ConfigCmd() {
	config.ReUnmarshalConfig()
}
