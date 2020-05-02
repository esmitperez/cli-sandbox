package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	// make our program aware about this command
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays this program's version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli-sandbox 0.1")
	},
}
