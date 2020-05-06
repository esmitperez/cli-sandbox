package cmd

import (
	"github.com/spf13/cobra"
)

var Version = "0.1"
var CmdName = "cli-sandbox"

var rootCmd = &cobra.Command{
	Use:     CmdName,
	Short:   "Esmit's attempt to make sense of Cobra",
	Version: Version,
}

// Execute is the entry point to our app
// If the execution of the root command fails we `panic`
func Execute() {
	// see https://golang.org/pkg/text/template/ for more on template customizations
	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}

Author: Esmit Perez - @mitiwifi
👓💻🔭🔬🌎🐕🐈🎸♏, ☀🌊🚣🌋🌋🌋🌊🚣
`)
	if err := rootCmd.Execute(); err != nil {
		panic("Something's wrong") // TODO introduce proper logging framework
	}
}
