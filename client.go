package smarthr

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"runtime"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Token      string
}

var userAgent = fmt.Sprintf("SmartHRGoClient/%s (%s)", version, runtime.Version())

func NewClient(c *Config) (*Client, error) {
	parsedURL, err := c.URL()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Client{
		URL:        parsedURL,
		HTTPClient: &http.Client{},
		Token:      c.Token,
	}, nil
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(out)
	return errors.WithStack(err)
}

func (c *Client) url(endpoint string) string {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, endpoint)
	return u.String()
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("User-Agent", userAgent)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return res, errors.WithStack(err)
	}
	if res.StatusCode < http.StatusBadRequest {
		return res, nil
	}
	var e Error
	if err := decodeBody(res, &e); err != nil {
		return res, nil
	}
	return res, errors.WithStack(e)
}

func (c *Client) get(ctx context.Context, endpoint string, opt interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.url(endpoint), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	v, err := query.Values(opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.URL.RawQuery = v.Encode()
	res, err := c.do(ctx, req)
	return res, errors.WithStack(err)
}

func (c *Client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, c.url(endpoint), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, err := c.do(ctx, req)
	return res, errors.WithStack(err)
}

func (c *Client) put(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, c.url(endpoint), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, err := c.do(ctx, req)
	return res, errors.WithStack(err)
}

func (c *Client) patch(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPatch, c.url(endpoint), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, err := c.do(ctx, req)
	return res, errors.WithStack(err)
}

func (c *Client) delete(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, c.url(endpoint), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.do(ctx, req)
	return res, errors.WithStack(err)
}
