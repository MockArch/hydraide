package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hydraidectl",
	Short: "HydrAIDE Control CLI",
	Long: `
💠 HydrAIDE Control CLI

Welcome to hydraidectl – your tool to install, restart, destroy and inspect your HydrAIDE system.

Usage:
  hydraidectl <command>

Try:
  hydraidectl init
  hydraidectl doctor
  hydraidectl restart
  hydraidectl destroy
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}
