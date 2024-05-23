package system

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	VERSION string = "0.1.0"
)

func GetVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display CLI version",
		Long:  `Print out project development version.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Infrasights CLI v%s\n", VERSION)
		},
	}
}
