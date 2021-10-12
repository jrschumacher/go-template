package cmd

import (
	"encoding/json"
	"os"

	"github.com/jrschumacher/go-template/internal/store"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	storeGetCmd = &cobra.Command{
		Use:   "get <id>",
		Short: "Get item to store",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			store := store.OpenStore(storeLocation)
			defer store.Close()

			id := args[0]

			result, err := store.Get(id)
			if err != nil {
				logrus.WithError(err).WithField("id", id).Fatal("Could not find item.")
			}

			json.NewEncoder(os.Stdin).Encode(result)
		},
	}
)

func init() {
	storeCmd.AddCommand(storeGetCmd)
}
