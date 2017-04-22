package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// parseCmd represents the version command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parses Tectonic-Installer config variable file",
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("Tectonic-Offline Version: v0.0.1")
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)
}