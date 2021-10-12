package cmd

import (
	"encoding/json"
	"os"

	"github.com/jrschumacher/go-template/internal/store"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	storeAddCmd = &cobra.Command{
		Use:   "add <json-data>",
		Short: "Add item to store",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			store := store.OpenStore(storeLocation)
			defer store.Close()

			var data map[string]interface{}
			if err := json.Unmarshal([]byte(args[0]), &data); err != nil {
				logrus.Fatal("Expects json object as string")
			}

			result, err := store.Add(data)
			if err != nil {
				logrus.WithError(err).Fatal("Unexpected error")
			}

			json.NewEncoder(os.Stdin).Encode(result)
		},
	}
)

func init() {
	storeCmd.AddCommand(storeAddCmd)
}
