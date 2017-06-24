package hatena

import (
	"fmt"
	"net/url"
	"regexp"
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

//bookmarkCommentUrl format : http://b.hatena.ne.jp/{userId}/{YYYYMMDD}#bookmark-{eid}
func CommentStar(bookmarkCommentUrl string) (*CommentStars, error) {
	return DefaultClient.CommentStar(bookmarkCommentUrl)
}

func (c *Client) CommentStar(bookmarkCommentUrl string) (*CommentStars, error) {

	//Validation URL format
	matched, err := regexp.MatchString("http:\\/\\/b\\.hatena\\.ne\\.jp\\/(.*)\\/[0-9]{8}#bookmark-[0-9]*", bookmarkCommentUrl)
	if !matched {
		return nil, err
	}

	val := url.Values{}
	val.Add("uri", bookmarkCommentUrl)

	req := starURL + "entry.json?" + val.Encode()
	fmt.Println(req)

	cs := &CommentStars{}
	err = c.get(req, cs, "json")
	fmt.Println(cs)

	return cs, err
}
