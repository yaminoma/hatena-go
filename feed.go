package hatena

import (
	"net/url"
)

const (
	feedURL = "http://b.hatena.ne.jp/"
)

type Entry struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Date          string `xml:"date"`
	Subject       string `xml:"subject"`
	Bookmarkcount int    `xml:"bookmarkcount"`
}

type Entries struct {
	Title       string  `xml:"channel>title"`
	Link        string  `xml:"channel>link"`
	Description string  `xml:"channel>description"`
	Entry       []Entry `xml:"item"`
}

func HotEntryAll() (*Entries, error) {
	return DefaultClient.HotEntryAll()
}

func (c *Client) HotEntryAll() (*Entries, error) {

	req := feedURL + "hotentry?mode=rss"

	e := &Entries{}
	err := c.get(req, e, "xml")

	return e, err
}

// Categories : social, economics, life, knowledge, it, fun, entertainment, game
func HotEntry(category string) (*Entries, error) {
	return DefaultClient.HotEntry(category)
}

func (c *Client) HotEntry(category string) (*Entries, error) {

	req := feedURL + "hotentry/" + category + ".rss"

	e := &Entries{}
	err := c.get(req, e, "xml")

	return e, err
}

func NewEntryAll(options map[string]string) (*Entries, error) {
	return DefaultClient.NewEntryAll(options)
}

func (c *Client) NewEntryAll(options map[string]string) (*Entries, error) {

	val := make(url.Values)
	val.Add("mode", "rss")
	for k, v := range options {
		val.Add(k, v)
	}

	req := feedURL + "entrylist?" + val.Encode()

	e := &Entries{}
	err := c.get(req, e, "xml")

	return e, err
}

// Categories : social, economics, life, knowledge, it, fun, entertainment, game
// Option : sort=recent, sort=popular, threshold={3}
func NewEntry(category string, options map[string]string) (*Entries, error) {
	return DefaultClient.NewEntry(category, options)
}

func (c *Client) NewEntry(category string, options map[string]string) (*Entries, error) {

	val := make(url.Values)
	for k, v := range options {
		val.Add(k, v)
	}
	req := feedURL + "entrylist/" + category + ".rss?" + val.Encode()

	e := &Entries{}
	err := c.get(req, e, "xml")

	return e, err
}

// searchType : keyword, title, tag
// Option : sort=recent, sort=popular, threshold={10}, date_begin={YYYY-MM-DD}, date_end={YYYY-MM-DD}, safe={on/off}
func SearchEntry(searchWord string, searchType string, options map[string]string) (*Entries, error) {
	return DefaultClient.SearchEntry(searchWord, searchType, options)
}

func (c *Client) SearchEntry(searchWord string, searchType string, options map[string]string) (*Entries, error) {

	val := make(url.Values)
	val.Add("q", searchWord)
	val.Add("mode", "rss")

	for k, v := range options {
		val.Add(k, v)
	}

	req := feedURL + "search/" + searchType + "?" + val.Encode()

	e := &Entries{}
	err := c.get(req, e, "xml")

	return e, err
}

// Option : sort=count, sort=eid, sort=recent
func SearchUrlEntry(searchUrl string, options map[string]string) (*Entries, error) {
	return DefaultClient.SearchUrlEntry(searchUrl, options)
}

func (c *Client) SearchUrlEntry(searchUrl string, options map[string]string) (*Entries, error) {

	val := make(url.Values)
	val.Add("url", searchUrl)
	val.Add("mode", "rss")

	for k, v := range options {
		val.Add(k, v)
	}

	req := feedURL + "entrylist?" + val.Encode()

	e := &Entries{}
	err := c.get(req, e, "xml")

	return e, err
}
