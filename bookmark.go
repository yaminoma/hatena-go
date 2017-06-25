package hatena

import (
	"fmt"
	_ "net/url"
)

var (
	bookmarkURL = fmt.Sprintf("http://api.b.hatena.ne.jp/%d/my/bookmark", version)
)

type BookmarkResponse struct {
	comment          string
	created_datetime string
	created_epoch    int
	user             string
	permalink        string
	private          bool
	tags             string
}

//func Get(urlStr string) (response string, err error) {
//	return oauthClient.GetBookmark(urlStr)
//}
//
//func (a *Authenticator) GetBookmark(urlStr string) (response string, err error) {
//
//	v := url.Values{}
//	v.Set("url", urlStr)
//	req := bookmarkURL + "?" + v.Encode()
//
//	b := new(string)
//	err := c.get(req, b, "json")
//
//	return b, errs
//}

//func BookmarkAdd(urlStr string, comment string, tags []string) (result string, err error) {
//	return Authenticator.Add(urlStr, comment, tags)
//}
//
//func (a *Authenticator) BookmarkAdd(urlStr string, comment string, tags []string) (result string, err error) {
//
//	v := url.Values{}
//	v.Set("url", urlStr)
//	req := countURL + "?" + v.Encode()
//
//	b := new(string)
//	err := c.get(req, b, "json")
//
//	return b, errs
//}
