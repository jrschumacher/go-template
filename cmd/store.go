package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	storeCmd = &cobra.Command{
		Use:   "store",
		Short: "Store subcommand",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}
)

func init() {
	rootCmd.AddCommand(storeCmd)
}
