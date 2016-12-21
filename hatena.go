package hatena

import (
	"fmt"
	_ "log"
	"net/http"
	_ "net/url"

	_ "github.com/markbates/goth"
)

const (
	acceptHeader = "application/json"
	version      = 1
)

var (
	defaultBaseURL = fmt.Sprintf("http://api.b.hatena.ne.jp/%d", version)

	// DefaultClient is the default client that is used by the wrapper functions
	// that don't require authorization.  If you need to authenticate, create
	// your own client with `Authenticator.NewClient`.
	DefaultClient = &Client{
		http: new(http.Client),
	}
)

//type Client struct {
//	Provider          goth.Provider
//	AccessToken       string
//	AccessTokenSecret string
//	BaseURL           *url.URL
//	UserAgent         string
//	BookMark          *Bookmark
//}

// Client is a client for working with the Hatena Web API.
// To create an authenticated client, use the
// `Authenticator.NewClient` method.  If you don't need to
// authenticate, you can use `DefaultClient`.
type Client struct {
	http *http.Client
}

//func NewHatena(provider goth.Provider, accessToken string, accessTokenSecret string) *Client {
//
//	//baseURL, _ := url.Parse(defaultBaseURL)
//
//	//cli := &Client{
//	//	Provider:          provider,
//	//	AccessToken:       accessToken,
//	//	AccessTokenSecret: accessTokenSecret,
//	//	BaseURL:           baseURL,
//	//}
//
//	////cli.Entry = &Entry{client: cli}
//	//cli.BookMark = &Bookmark{client: cli}
//	////cli.Tag = &Tag{client: cli}
//	////cli.User = &User{client: cli}
//
//	return
//}
