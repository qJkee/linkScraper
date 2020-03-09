package cli

import (
	"errors"
	"fmt"

	"github.com/qJkee/linkScraper/api"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "stop crawling website",
	Example: "stop www.example.com",
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
		resp, err := cl.StopCrawl(cmd.Context(), &api.URLRequest{Url: args[0]})
		if err != nil {
			return err
		}
		fmt.Println(resp.ServerResp)
		return nil
	},
}
