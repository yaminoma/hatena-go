package hatena

import (
	"net/url"

	"github.com/parnurzeal/gorequest"
)

const (
	countURL = "http://api.b.st-hatena.com/entry.count"
)

func Count(query string) (body string, errs []error) {
	return DefaultClient.Count(query)
}

func (c *Client) Count(urlStr string) (body string, errs []error) {

	v := url.Values{}
	v.Set("url", urlStr)
	req := countURL + "?" + v.Encode()

	request := gorequest.New()
	resp, body, errs := request.Get(req).End()

	if resp.StatusCode != 200 {
		return body, errs
	}
	return body, errs
}
