package main

import (
	"fmt"
	_ "log"
	"net/http"
	"net/url"

	"github.com/8pockets/hatena-go"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Hatena's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"
const ConsumerKey = "***************"
const ConsumerSecret = "***************"

var (
	scopes = []string{hatena.HatenaReadPrivate}
	//scopes = []string{hatena.HatenaReadPublic, hatena.HatenaReadPrivate}
	auth = hatena.NewAuthenticator(ConsumerKey, ConsumerSecret, redirectURI, scopes)
)

func main() {
	// first start an HTTP server
	http.HandleFunc("/auth", authUrl)

	http.HandleFunc("/callback", token)
	http.Handle("/profile", &hatena.AuthHandler{Handler: profile})

	http.ListenAndServe(":8080", nil)

	//response, err := hatena.Get("https://api.nasa.gov/")
}

func authUrl(w http.ResponseWriter, r *http.Request) {
	url := auth.AuthURL(w, r)
	fmt.Fprintf(w, "<a href='%s'>Auth</a>", url)
	fmt.Printf("1. Go to %s\n2. Authorize the application\n", url)
}

func token(w http.ResponseWriter, r *http.Request) {
	credentials := auth.Token(w, r)
	fmt.Printf("Token: %s, Secret: %s", credentials.Token, credentials.Secret)

	http.Redirect(w, r, "/profile", 302)
}

func profile(w http.ResponseWriter, r *http.Request) {
	var dms []map[string]interface{}

	form := url.Values{}
	form.Add("url", "http://www.fashionsnap.com/the-posts/2017-01-04/rap-year-book/")

	if err := auth.ApiPost(
		"http://api.b.hatena.ne.jp/1/my/bookmark",
		form,
		&dms); err != nil {
		http.Error(w, "Error getting api, "+err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "%s", dms)
}
