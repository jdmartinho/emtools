package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetVersionCmd(v string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of emtools",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("emtools version: %s", v)
		},
	}
}
