package hatena

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
)

const (
	version = 1
)

var (
	// DefaultClient is the default client that is used by the wrapper functions
	// that don't require authorization.  If you need to authenticate, create
	// your own client with `Authenticator.NewClient`.
	DefaultClient = &Client{
		http: new(http.Client),
	}
)

// Client is a client for working with the Hatena Web API.
// To create an authenticated client, use the
// `Authenticator.NewClient` method.  If you don't need to
// authenticate, you can use `DefaultClient`.
type Client struct {
	http *http.Client
}

func (c *Client) get(url string, result interface{}, format string) error {
	resp, err := c.http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("http status is not http.StatusOK")
	}

	if format == "xml" {
		decoder := xml.NewDecoder(resp.Body)
		decoder.Strict = false
		decoder.Decode(result)
	} else if format == "json" {
		err = json.NewDecoder(resp.Body).Decode(result)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) post(url string, data url.Values, result interface{}) error {
	resp, err := c.http.PostForm(url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("http status is not htp.StatusOK")
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) delete(url string) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	//削除成功の場合は204(http.StatusNoContent)が返却される
	if resp.StatusCode != http.StatusNoContent {
		return errors.New("http status is not http.StatusNoContent")
	}

	return nil
}
