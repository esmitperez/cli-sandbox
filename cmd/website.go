package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	// make our program aware about this command
	rootCmd.AddCommand(websiteCmd)
}

var websiteCmd = &cobra.Command{
	Use:   "website",
	Short: "Displays Esmit's personal site link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nhttps://esmit.me")
	},
}
