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

type CommentStars struct {
	Entries []struct {
		Stars []struct {
			Quote string `json:"quote"`
			Name  string `json:"name"`
		} `json:"stars"`
		CanComment int    `json:"can_comment"`
		URI        string `json:"uri"`
	} `json:"entries"`
	CanComment int `json:"can_comment"`
}

// はてなスターカウントAPI
// 「指定したブログのエントリに全部でいくつのスターがつけられているのか」という総数を取得できるAPI
func UserStar(username string) (*UserStars, error) {
	return DefaultClient.UserStar(username)
}

func (c *Client) UserStar(username string) (*UserStars, error) {

	uri := "http://b.hatena.ne.jp/" + username + "/"

	val := url.Values{}
	val.Add("uri", uri)

	req := starURL + "blog.json?" + val.Encode()

	us := &UserStars{}
	err := c.get(req, us, "json")

	return us, err
}

// はてなスター取得 API
// ある URL に対して付与されたスターを取得できる。
func GetStar(urlStr string) (*Star, error) {
	return DefaultClient.GetStar(urlStr)
}

func (c *Client) GetStar(urlStr string) (*Star, error) {

	val := url.Values{}
	val.Add("uri", urlStr)

	req := starURL + "entry.json?" + val.Encode()

	s := &Star{}
	err = c.get(req, s, "json")

	return s, err
}
