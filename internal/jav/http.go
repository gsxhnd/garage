package jav

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const defaultUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"

type httpClient struct {
	client *resty.Client
	delay  time.Duration
}

func newHTTPClient(cfg Config) (*httpClient, error) {
	client := resty.New()
	client.SetTimeout(cfg.Timeout)
	client.SetHeader("User-Agent", defaultUA)
	client.SetHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	client.SetHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	cookie := cfg.Cookie
	if cookie == "" {
		cookie = "over18=1; locale=zh"
	} else if !strings.Contains(cookie, "over18=") {
		cookie += "; over18=1"
	}
	client.SetHeader("Cookie", cookie)

	if cfg.Proxy != "" {
		u, err := url.Parse(cfg.Proxy)
		if err != nil || u.Scheme == "" || u.Host == "" {
			return nil, fmt.Errorf("jav: invalid proxy %q", cfg.Proxy)
		}
		client.SetProxy(cfg.Proxy)
	}

	return &httpClient{
		client: client,
		delay:  cfg.RandomDelay,
	}, nil
}

func (c *httpClient) Get(ctx context.Context, rawURL string, extra map[string]string) ([]byte, error) {
	if err := c.wait(ctx); err != nil {
		return nil, err
	}

	req := c.client.R().SetContext(ctx)
	for k, v := range extra {
		req.SetHeader(k, v)
	}

	resp, err := req.Get(rawURL)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {
	case http.StatusOK:
		return resp.Body(), nil
	case http.StatusNotFound:
		return nil, errNotFound
	default:
		if resp.StatusCode() >= 400 {
			return nil, fmt.Errorf("GET %s: status %d", rawURL, resp.StatusCode())
		}
		return resp.Body(), nil
	}
}

func (c *httpClient) wait(ctx context.Context) error {
	if c.delay <= 0 {
		return ctx.Err()
	}
	timer := time.NewTimer(c.delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
