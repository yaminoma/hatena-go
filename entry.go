package hatena_go

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

const (
	entryURL = "http://b.hatena.ne.jp/entry/json/"
)

type EntryInformation struct {
	Eid        int        `json:"eid"`
	Title      string     `json:"title"`
	Count      int        `json:"count"`
	Url        string     `json:"url"`
	Entry_url  string     `json:"entry_url"`
	Screenshot string     `json:"screenshot"`
	Bookmarks  []Bookmark `json:"bookmarks"`
	Related    []Related  `json:"related"`
}

type Bookmark struct {
	User      string   `json:"user"`
	Comment   string   `json:"comment"`
	Timestamp string   `json:"timestamp"`
	Tags      []string `json:"tags"`
}

type Related struct {
	Eid       int    `json:"eid"`
	Title     string `json:"title"`
	Count     int    `json:"count"`
	Url       string `json:"url"`
	Entry_url string `json:"entry_url"`
}

func EntryInfo(url string) (*EntryInformation, error) {
	return DefaultClient.EntryInfo(url)
}

func (c *Client) EntryInfo(url string) (*EntryInformation, error) {

	req := entryURL + "?url=" + url
	request := gorequest.New()
	resp, body, errs := request.Get(req).End()

	e := EntryInformation{}
	if resp.StatusCode != 200 {
		return &e, errs[0]
	}

	err := json.Unmarshal([]byte(body), &e)

	return &e, err
}
