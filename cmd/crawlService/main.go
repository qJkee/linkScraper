package main

import (
	"flag"
	"log"

	"github.com/qJkee/linkScraper/internal/crawler"
	"github.com/qJkee/linkScraper/internal/crawler/service"
)

var listenAddr = flag.String("addr", ":2020", "gRPC listen addr")

func main() {
	flag.Parse()
	srv := service.NewCrawler()
	cr := crawler.New(srv)
	log.Fatal(cr.Start(*listenAddr))
}
