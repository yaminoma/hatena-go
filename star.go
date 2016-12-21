package hatena

import (
	"encoding/json"
	"net/url"
	"regexp"

	"github.com/parnurzeal/gorequest"
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
	Uri       string `json:"uri"`
	StarCount int    `json:"star_count"`
}

type CommentStars struct {
	Entries []struct {
		Stars []struct {
			Quote string `json:"quote"`
			Name  string `json:"name"`
		} `json:"stars"`
		CanComment int    `json:"can_comment"`
		Uri        string `json:"uri"`
	} `json:"entries"`
	CanComment int `json:"can_comment"`
}

func UserStar(username string) (*UserStars, error) {
	return DefaultClient.UserStar(username)
}

func (c *Client) UserStar(username string) (*UserStars, error) {

	uri := "http://b.hatena.ne.jp/" + username + "/"

	val := url.Values{}
	val.Add("uri", uri)

	req := starURL + "blog.json?" + val.Encode()

	request := gorequest.New()
	resp, body, errs := request.Get(req).End()

	u := UserStars{}
	if resp.StatusCode != 200 {
		return &u, errs[0]
	}

	err := json.Unmarshal([]byte(body), &u)

	return &u, err
}

//bookmarkCommentUrl format : http://b.hatena.ne.jp/{userId}/{YYYYMMDD}#bookmark-{eid}
func CommentStar(bookmarkCommentUrl string) (*CommentStars, error) {
	return DefaultClient.CommentStar(bookmarkCommentUrl)
}

func (c *Client) CommentStar(bookmarkCommentUrl string) (*CommentStars, error) {

	//Validation URL format
	res, regErr := regexp.MatchString("http:\\/\\/b\\.hatena\\.ne\\.jp\\/(.*)\\/[0-9]{8}#bookmark-[0-9]*", bookmarkCommentUrl)
	if res != true {
		return &CommentStars{}, regErr.(error)
	}

	val := url.Values{}
	val.Add("uri", bookmarkCommentUrl)

	req := starURL + "entry.json?" + val.Encode()
	request := gorequest.New()
	resp, body, errs := request.Get(req).End()

	s := CommentStars{}
	if resp.StatusCode != 200 {
		return &s, errs[0]
	}

	err := json.Unmarshal([]byte(body), &s)

	return &s, err
}
