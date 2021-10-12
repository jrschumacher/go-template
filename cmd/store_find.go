package cmd

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/jrschumacher/go-template/internal/store"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	storeFindCmd = &cobra.Command{
		Use:   "find [<key=value>...]",
		Short: "Find items in store",
		Run: func(cmd *cobra.Command, args []string) {
			store := store.OpenStore(storeLocation)
			defer store.Close()

			params := make(map[string]interface{})
			for _, arg := range args {
				parts := strings.Split(arg, "=")
				if len(parts) <= 1 {
					logrus.Fatal("Find queries expects 'key=value' format")
				}
				// add to params
				params[parts[0]] = strings.Join(parts[1:], "=")
			}

			result, err := store.Find(params)
			if err != nil {
				logrus.WithError(err).WithFields(params).Fatal("Could not find items.")
			}

			json.NewEncoder(os.Stdin).Encode(result)
		},
	}
)

func init() {
	storeCmd.AddCommand(storeFindCmd)
}
