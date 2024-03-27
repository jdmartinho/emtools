package commands

import (
	"fmt"

	"github.com/jdmartinho/emtools/internal/emtools"
	"github.com/spf13/cobra"
)

func GetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "emtools",
		Version: emtools.Version,
		Short:   "Tools for Engineering Managers",
		Long:    `TBD`,
		Args:    cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("emtools is working")
		},
	}
	return cmd
}
