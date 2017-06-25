package hatena

import (
	"fmt"
	"net/url"
)

var (
	bookmarkURL      = fmt.Sprintf("http://api.b.hatena.ne.jp/%d/my/bookmark", version)
	bookmarkEntryURL = fmt.Sprintf("http://api.b.hatena.ne.jp/%d/entry", version)
)

type Bookmark struct {
	Favorites       []interface{} `json:"favorites"`
	CommentRaw      string        `json:"comment_raw"`
	Private         bool          `json:"private"`
	Eid             int           `json:"eid"`
	CreatedEpoch    int           `json:"created_epoch"`
	Tags            []interface{} `json:"tags"`
	Permalink       string        `json:"permalink"`
	Comment         string        `json:"comment"`
	CreatedDatetime string        `json:"created_datetime"`
	User            string        `json:"user"`
}

type BookmarkForm struct {
	uri           string
	comment       string
	tags          []string
	post_twitter  bool
	post_facebook bool
	post_mixi     bool
	post_evernote bool
	send_mail     bool
	private       bool
}

// ブックマーク API
// ブックマーク情報を取得する
func GetBookmark(uri string) (*Bookmark, error) {
	return oauthClient.GetBookmark(uri)
}

func (o *Authenticator) GetBookmark(uri string) (*Bookmark, error) {

	form := url.Values{}
	form.Set("url", uri)

	b := &Bookmark{}
	err := o.apiGet(bookmarkURL, form, b)

	return b, err
}

// ブックマーク API
// ブックマークを追加または更新する
func AddBookmark(uri string, br BookmarkForm) (*Bookmark, error) {
	return oauthClient.AddBookmark(uri, br)
}

func (o *Authenticator) AddBookmark(uri string, br BookmarkForm) (*Bookmark, error) {

	// BookmarkFormをurl.Valueに変換
	form := url.Values{}
	form.Set("url", uri)

	b := &Bookmark{}
	err := o.apiPost(bookmarkURL, form, b)

	return b, err
}

// ブックマーク API
// ブックマークを削除する
func DeleteBookmark(uri string) error {
	return oauthClient.DeleteBookmark(uri)
}

func (o *Authenticator) DeleteBookmark(uri string) error {

	form := url.Values{}
	form.Set("url", uri)

	b := &Bookmark{}
	err := o.apiDelete(bookmarkURL, form, b)

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
func GetBookmarkedEntry(uri string) (*BookmarkEntry, error) {
	return oauthClient.GetBookmarkedEntry(uri)
}

func (o *Authenticator) GetBookmarkedEntry(uri string) (*BookmarkEntry, error) {

	form := url.Values{}
	form.Set("url", uri)

	b := &BookmarkEntry{}
	err := o.apiGet(bookmarkEntryURL, form, b)

	return b, err
}
