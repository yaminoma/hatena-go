package hatena

import (
	_ "fmt"
	_ "net/url"

	_ "github.com/mrjones/oauth"
	_ "github.com/parnurzeal/gorequest"
)

const (
	bookmarkURL = "http://api.b.st-hatena.com/entry.count"
)

type Bookmark struct {
	client *Client
}

type BookmarkResponse struct {
	comment          string
	created_datetime string
	created_epoch    int
	user             string
	permalink        string
	private          bool
	tags             string
}

func (s *Bookmark) Add(urlStr string, comment string, tags []string) (response string, err error) {

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

func (s *Bookmark) Delete(url string) (err error) {
	return
}
