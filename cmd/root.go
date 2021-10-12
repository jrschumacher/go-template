package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	storeLocation string
	rootCmd       = &cobra.Command{
		Use:   "app [command]",
		Short: "go template",
		Long:  `A go template app`,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&storeLocation, "store-path", "s", "./tmp-store.db", "Set the store location")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
