package main

import (
	"fmt"
	"os"

	"github.com/jdmartinho/emtools/cmd/commands"
	"github.com/jdmartinho/emtools/internal/emtools"
)

func main() {
	initCli()
}

func initCli() {
	rootCmd := commands.GetRootCmd()
	rootCmd.AddCommand(commands.GetVersionCmd(emtools.Version))

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
