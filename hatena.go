package hatena_go

import (
	"fmt"
	_ "log"
	"net/url"

	"github.com/markbates/goth"
)

const (
	acceptHeader = "application/json"
	version      = 1
)

var (
	defaultBaseURL = fmt.Sprintf("http://api.b.hatena.ne.jp/%d", version)
)

type Client struct {
	Provider          goth.Provider
	AccessToken       string
	AccessTokenSecret string
	BaseURL           *url.URL
	UserAgent         string
	BookMark          *Bookmark
	//Entry     *Entry
	//Tag       *Tag
	//User      *User
}

func NewHatena(provider goth.Provider, accessToken string, accessTokenSecret string) *Client {

	baseURL, _ := url.Parse(defaultBaseURL)

	cli := &Client{
		Provider:          provider,
		AccessToken:       accessToken,
		AccessTokenSecret: accessTokenSecret,
		BaseURL:           baseURL,
	}

	//cli.Entry = &Entry{client: cli}
	cli.BookMark = &Bookmark{client: cli}
	//cli.Tag = &Tag{client: cli}
	//cli.User = &User{client: cli}

	return cli
}
