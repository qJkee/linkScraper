package service

import (
	"bytes"
	"errors"
	"net/url"
	"sync"
)

type Crawler struct {
	workers *sync.Map
	bufPool *sync.Pool
}

func NewCrawler() *Crawler {
	return &Crawler{
		workers: new(sync.Map),
		bufPool: &sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}

func (c *Crawler) ProcessURL(link string, parallelism int64) error {
	targetURL, err := url.Parse(link)
	if err != nil {
		return err
	}
	if targetURL.Scheme == "" {
		return errors.New("No protocol scheme specified ")
	}
	w := newWorker(targetURL, parallelism)
	_, ok := c.workers.LoadOrStore(link, w)
	if ok {
		return errors.New("Link already parsing ")
	}
	go w.start()
	return nil
}

func (c *Crawler) StopProcessing(link string) error {
	w, ok := c.workers.Load(link)
	if !ok {
		return errors.New("URL is not parsing ")
	}
	w.(*Worker).stop()
	return nil
}

func (c *Crawler) List() string {
	buf := c.bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		c.bufPool.Put(buf)
	}()
	c.workers.Range(func(key, value interface{}) bool {
		buf.WriteString("----------------------------- \n")
		buf.WriteString(value.(*Worker).tree.Print())
		buf.WriteString("\n")
		return true
	})
	return buf.String()
}
