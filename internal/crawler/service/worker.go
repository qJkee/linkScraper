package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"

	gotree "github.com/DiSiqueira/GoTree"

	"github.com/PuerkitoBio/goquery"
)

type Worker struct {
	rootURL *url.URL
	httpCl  *http.Client
	tree    gotree.Tree
	ctx     context.Context
	cancel  context.CancelFunc
	uniqMap *sync.Map
	sem     *semaphore.Weighted
}

func newWorker(link *url.URL, maxParallel int64) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		rootURL: link,
		httpCl:  &http.Client{Timeout: 10 * time.Second},
		tree:    gotree.New(link.String()),
		ctx:     ctx,
		cancel:  cancel,
		uniqMap: new(sync.Map),
		sem:     semaphore.NewWeighted(maxParallel),
	}
}

func (w *Worker) start() {
	w.uniqMap.Store(w.rootURL.String(), struct{}{})
	w.process(w.ctx, w.rootURL, w.tree)
}

func (w *Worker) stop() {
	w.cancel()
}

func (w *Worker) process(ctx context.Context, url *url.URL, tree gotree.Tree) {
	err := w.sem.Acquire(ctx, 1)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		fmt.Println("Can't acquire from semaphore ", err)
		return
	}
	defer w.sem.Release(1)
	body, err := w.getHtml(ctx, url)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		fmt.Println("Can't get html for url:", url.String(), " error: ", err)
		return
	}
	urls, err := parseHTML(body)
	if err != nil {
		fmt.Println("Can't parse html,err: ", err)
	}
	body.Close()
	for _, v := range urls {
		if strings.HasPrefix(v, "/") {
			url.Path = v
			v = url.String()
		} else if strings.HasPrefix(v, "#") {
			v = url.String() + v
		}
		_, ok := w.uniqMap.LoadOrStore(v, struct{}{})
		if ok {
			continue
		}
		u, err := url.Parse(v)
		if err != nil {
			fmt.Println("Can't parse found url,err: ", err)
			continue
		}
		if u.Host != w.rootURL.Host {
			continue
		}
		newTree := tree.Add(v)
		go w.process(ctx, u, newTree)
	}
}

func parseHTML(data io.Reader) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		link, exist := selection.Attr("href")
		if !exist {
			return
		}
		result = append(result, link)
	})
	return result, nil
}

func (w *Worker) getHtml(ctx context.Context, link *url.URL) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", link.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := w.httpCl.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors.New("Invalid response code " + resp.Status)
	}
	return resp.Body, nil
}
