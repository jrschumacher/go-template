package cmd

import (
	"encoding/json"
	"os"

	"github.com/jrschumacher/go-template/internal/store"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	storeUpdateCmd = &cobra.Command{
		Use:   "update <id> <json-data>",
		Short: "Update item in store",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			store := store.OpenStore(storeLocation)
			defer store.Close()

			id := args[0]
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(args[1]), &data); err != nil {
				logrus.Fatal("Expects json object as string")
			}

			result, err := store.Update(id, data)
			if err != nil {
				logrus.WithError(err).WithField("id", id).Fatal("Could not update item.")
			}

			json.NewEncoder(os.Stdin).Encode(result)
		},
	}
)

func init() {
	storeCmd.AddCommand(storeUpdateCmd)
}
