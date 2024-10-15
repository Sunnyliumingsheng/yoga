package cmd

import (
	"github.com/spf13/cobra"

	"cli/config"
)

var rootCmd = &cobra.Command{
	"Use":config.Config.
}
