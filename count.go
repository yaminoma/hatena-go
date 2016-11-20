package hatena_go

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

func (c *Client) Add(urlStr string, comment string, tags []string) (response string, err error) {

	////u, _ := url.Parse(urlStr)
	//tag := strings.Join(tags[:], ",")
	//params := map[string]string{"url": urlStr, "comment": comment, "tag": tag}

	//response, err = s.client.Provider.consumer.Get(
	//	s.client.defaultBaseURL,
	//	params,
	//	s.client.AccessToken)

	//if err != nil {
	//	return
	//}

	//defer response.Body.Close()

	return
}

func (c *Client) Delete(url string) (err error) {
	return
}
