package hatena

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

var (
	bookmarkURL      = fmt.Sprintf("http://api.b.hatena.ne.jp/%d/my/bookmark", version)
	bookmarkEntryURL = fmt.Sprintf("http://api.b.hatena.ne.jp/%d/entry", version)
)

type Bookmark struct {
	CommentRaw      string   `json:"comment_raw"`
	Private         bool     `json:"private"`
	Eid             int      `json:"eid"`
	CreatedEpoch    int      `json:"created_epoch"`
	Tags            []string `json:"tags"`
	Permalink       string   `json:"permalink"`
	Comment         string   `json:"comment"`
	CreatedDatetime string   `json:"created_datetime"`
	User            string   `json:"user"`
}

type BookmarkForm struct {
	uri           string   `json:"uri"`
	comment       string   `json:"comment"`
	tags          []string `json:"tags"`
	post_twitter  bool     `json:"post_twitter"`
	post_facebook bool     `json:"post_facebook"`
	post_mixi     bool     `json:"post_mixi"`
	post_evernote bool     `json:"post_evernote"`
	send_mail     bool     `json:"send_mail"`
	private       bool     `json:"private"`
}

// ブックマーク API
// ブックマーク情報を取得する
func (a *Authenticator) GetBookmark(uri string) (*Bookmark, error) {

	values := make(url.Values)
	values.Set("url", uri)

	b := &Bookmark{}
	err := a.apiGet(bookmarkURL, values, b)

	return b, err
}

// ブックマーク API
// ブックマークを追加または更新する
func (a *Authenticator) AddBookmark(br BookmarkForm) (*Bookmark, error) {

	values, err := query.Values(br)
	if err != nil {
		return nil, err
	}

	b := &Bookmark{}
	err = a.apiPost(bookmarkURL, values, b)

	return b, err
}

// ブックマーク API
// ブックマークを削除する
func (a *Authenticator) DeleteBookmark(uri string) error {

	values := make(url.Values)
	values.Set("url", uri)

	b := &Bookmark{}
	err := a.apiDelete(bookmarkURL, values, b)

	return err
}

type BookmarkEntry struct {
	Count                 int    `json:"count"`
	FaviconURL            string `json:"favicon_url"`
	Eid                   int    `json:"eid"`
	ImageLastEditor       string `json:"image_last_editor"`
	EntryURL              string `json:"entry_url"`
	RootURL               string `json:"root_url"`
	IsInvalidURL          bool   `json:"is_invalid_url"`
	SmartphoneAppEntryURL string `json:"smartphone_app_entry_url"`
	TitleLastEditor       string `json:"title_last_editor"`
	ImageURL              string `json:"image_url"`
	URL                   string `json:"url"`
	Title                 string `json:"title"`
	HasAsin               bool   `json:"has_asin"`
	ImageHatenaURL        string `json:"image_hatena_url"`
}

// エントリー API
// ブックマークされたエントリーの情報を取得する
func (a *Authenticator) GetBookmarkedEntry(uri string) (*BookmarkEntry, error) {

	values := make(url.Values)
	values.Set("url", uri)

	b := &BookmarkEntry{}
	err := a.apiGet(bookmarkEntryURL, values, b)

	return b, err
}
