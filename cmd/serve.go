package cmd

import (
	"github.com/jrschumacher/go-template/api"
	"github.com/spf13/cobra"
)

var (
	host     string
	port     int
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve the Restful API",
		Run: func(cmd *cobra.Command, args []string) {
			params := api.NewServeParams()
			params.StoreLocation = storeLocation
			params.Host = host
			params.Port = port

			api.Serve(params)
		},
	}
)

func init() {
	serveCmd.Flags().StringVar(&host, "host", "", "Server will listen to requests to this hostname")
	serveCmd.Flags().IntVar(&port, "port", 8080, "Server will listen to requests to this port")
	rootCmd.AddCommand(serveCmd)
}
