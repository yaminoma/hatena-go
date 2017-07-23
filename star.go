package hatena

import (
	"net/url"
)

const (
	starURL = "http://s.hatena.ne.jp/"
)

type UserStars struct {
	Count struct {
		Green  int    `json:"green"`
		Blue   int    `json:"blue"`
		Red    int    `json:"red"`
		Purple int    `json:"purple"`
		Yellow string `json:"yellow"`
	} `json:"count"`
	Title     string `json:"title"`
	URI       string `json:"uri"`
	StarCount int    `json:"star_count"`
}

type EntryStars struct {
	Entries    []EntryPost `json:"entries"`
	CanComment int         `json:"can_comment"`
}
type EntryPost struct {
	Stars        []Star         `json:"stars"`
	CanComment   int            `json:"can_comment"`
	ColoredStars []ColoredStars `json:"colored_stars"`
	URI          string         `json:"uri"`
}
type Star struct {
	Count int    `json:"count"`
	Quote string `json:"quote"`
	Name  string `json:"name"`
}
type ColoredStars struct {
	Stars []Star `json:"stars"`
	Color string `json:"color"`
}

// はてなスターカウントAPI
// 「指定したブログのエントリに全部でいくつのスターがつけられているのか」という総数を取得できるAPI
func UserStar(username string) (*UserStars, error) {
	return DefaultClient.UserStar(username)
}

func (c *Client) UserStar(username string) (*UserStars, error) {

	uri := "http://b.hatena.ne.jp/" + username + "/"

	v := make(url.Values)
	v.Add("uri", uri)

	req := starURL + "blog.json?" + v.Encode()

	us := &UserStars{}
	err := c.get(req, us, "json")

	return us, err
}

// はてなスター取得 API
// ある URL に対して付与されたスターを取得できる。
func GetStar(urlStr string) (*EntryStars, error) {
	return DefaultClient.GetStar(urlStr)
}

func (c *Client) GetStar(urlStr string) (*EntryStars, error) {

	v := make(url.Values)
	v.Add("uri", urlStr)

	req := starURL + "entry.json?" + v.Encode()

	e := &EntryStars{}
	err := c.get(req, e, "json")

	return e, err
}
