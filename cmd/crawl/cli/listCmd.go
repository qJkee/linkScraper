package cli

import (
	"fmt"

	"github.com/qJkee/linkScraper/api"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list of all crawling websites",
	Example: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		cl, err := NewClient(rpcAddr)
		if err != nil {
			return err
		}
		resp, err := cl.CrawlList(cmd.Context(), &api.Empty{})
		if err != nil {
			return err
		}
		fmt.Println(string(resp.SiteTree))
		return nil
	},
}
