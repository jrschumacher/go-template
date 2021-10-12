package cmd

import (
	"encoding/json"
	"os"

	"github.com/jrschumacher/go-template/internal"
	"github.com/spf13/cobra"
)

var (
	longVersion bool
	versionCmd  = &cobra.Command{
		Use:   "version",
		Short: "Print app version",
		Run: func(cmd *cobra.Command, args []string) {
			json.NewEncoder(os.Stdout).Encode(internal.GetVersion())
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
