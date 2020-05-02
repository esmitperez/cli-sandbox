package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-sandbox",
	Short: "Esmit's attempt to make sense of Cobra",
}

// Execute is the entry point to our app
// If the execution of the root command fails we `panic`
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic("Something's wrong") // TODO introduce proper logging framework
	}
}
