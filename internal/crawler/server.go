package crawler

import (
	"net"

	"github.com/qJkee/linkScraper/api"
	"github.com/qJkee/linkScraper/internal/crawler/service"

	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

type Server struct {
	srv *service.Crawler
}

func New(service *service.Crawler) *Server {
	return &Server{srv: service}
}

func (s *Server) Start(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	grpcServer := grpc.NewServer()
	api.RegisterCrawlerServer(grpcServer, s)
	return grpcServer.Serve(ln)
}

func (s *Server) StartCrawl(_ context.Context, req *api.URLRequest) (*api.StatusMessage, error) {
	err := s.srv.ProcessURL(req.Url, req.Parallelism)
	if err != nil {
		return nil, err
	}
	return &api.StatusMessage{ServerResp: "Starting to process url " + req.Url}, nil
}

func (s *Server) StopCrawl(_ context.Context, req *api.URLRequest) (*api.StatusMessage, error) {
	err := s.srv.StopProcessing(req.Url)
	if err != nil {
		return nil, err
	}
	return &api.StatusMessage{ServerResp: "Stopping to process url " + req.Url}, nil
}

func (s *Server) CrawlList(context.Context, *api.Empty) (*api.CrawlData, error) {
	return &api.CrawlData{SiteTree: []byte(s.srv.List())}, nil
}
