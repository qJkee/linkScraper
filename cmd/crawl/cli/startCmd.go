package cli

import (
	"errors"
	"fmt"

	"github.com/qJkee/linkScraper/api"

	"github.com/spf13/cobra"
)

var parallelism int64

func init() {
	startCmd.Flags().Int64Var(&parallelism, "parallelism", 5, "maximum gorutines to use for scrape")
}

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "start crawling website",
	Example: "start www.example.com",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("No target url provided ")
		} else if len(args) > 1 {
			return errors.New("Please provide only 1 url ")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := NewClient(rpcAddr)
		if err != nil {
			return err
		}
		resp, err := cl.StartCrawl(cmd.Context(), &api.URLRequest{Url: args[0], Parallelism: parallelism})
		if err != nil {
			return err
		}
		fmt.Println(resp.ServerResp)
		return nil
	},
}
