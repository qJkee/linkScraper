package cli

import (
	"log"

	"github.com/spf13/cobra"
)

const defaultRpcAddr = "localhost:2020"

var rpcAddr string

func init() {
	rootCmd.AddCommand(listCmd, startCmd, stopCmd)
	rootCmd.Flags().StringVar(&rpcAddr, "server", defaultRpcAddr, "grpc address of server")
}

var rootCmd = &cobra.Command{
	Use:   "crawl",
	Short: "GRPC CLI for crawler service",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
