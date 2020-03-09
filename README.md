# Simple scraping tool

Tool used to scrape links from the given URL.


## Start a server

* Build a server with go build(run `go build` inside `cmd/crawlService` dir)
* Start with `./crawlService`
* Optionally, you can specify GRPC server address by providing `--addr` flag. EXAMPLE: `--addr localhost:5000`


## Work with server

* This software ships with CLI to work with client.
* Build it with go build(run `go build` inside `cmd/crawl` dir) or `go install` to install it to the system 
* It's self documented by Cobra, but there are some examples

`crawl list` - shows current websites three map

`crawl start https://google.com` - start crawling `https://google.com`

`crawl stop https://google.com` - stops crawling `https://google.com`

To specify service GRPC address(if you are using custom one with flag), use `--server` flag for each cli call

To specify parallelism for parsing url, use `--parallelism` flag when calling `crawl start`