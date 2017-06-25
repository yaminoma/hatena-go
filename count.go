// はてなブックマーク件数取得API
package hatena

import (
	"net/url"
)

const (
	countURL = "http://api.b.st-hatena.com/entry.count"
)

func Count(query string) (int, error) {
	return DefaultClient.Count(query)
}

func (c *Client) Count(urlStr string) (int, error) {

	v := url.Values{}
	v.Set("url", urlStr)
	req := countURL + "?" + v.Encode()

	co := new(int)
	err := c.get(req, co, "json")

	return *co, err

}
