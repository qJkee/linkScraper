package cli

import (
	"context"

	"github.com/qJkee/linkScraper/api"

	"google.golang.org/grpc"
)

type Client struct {
	cl api.CrawlerClient
}

func (c *Client) StartCrawl(ctx context.Context, in *api.URLRequest, opts ...grpc.CallOption) (*api.StatusMessage, error) {
	return c.cl.StartCrawl(ctx, in, opts...)
}

func (c *Client) StopCrawl(ctx context.Context, in *api.URLRequest, opts ...grpc.CallOption) (*api.StatusMessage, error) {
	return c.cl.StopCrawl(ctx, in, opts...)
}

func (c *Client) CrawlList(ctx context.Context, in *api.Empty, opts ...grpc.CallOption) (*api.CrawlData, error) {
	return c.cl.CrawlList(ctx, in, opts...)
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		cl: api.NewCrawlerClient(conn),
	}, nil
}
