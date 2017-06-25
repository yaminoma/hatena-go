package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/8pockets/hatena-go"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Hatena's developer portal
// and enter this value.
const ConsumerKey = "2wDjTLiLsfRplA=="
const ConsumerSecret = "ts+bQgDIp/GI1I5q+v8Ca+12pA0="
const redirectURI = "http://localhost:8080/callback"

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
}

func authUrl(w http.ResponseWriter, r *http.Request) {
	url, err := auth.AuthURL(w, r)
	if err != nil {
		http.Error(w, "Error getting temp cred, "+err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "<a href='%s'>Authentication</a>", url)
}

func token(w http.ResponseWriter, r *http.Request) {
	credentials, err := auth.Token(w, r)
	if err != nil {
		http.Error(w, "Error getting temp cred, "+err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/profile", 302)
}

func profile(w http.ResponseWriter, r *http.Request) {

	prof, _ := auth.GetProfile()
	json.NewEncoder(w).Encode(prof)

	//resp, err := auth.GetBookmarkedEntry("http://8pockets.hatenablog.com/entry/2013/12/30/162516")
	//if err != nil {
	//	http.Error(w, "ERROR: "+err.Error(), 500)
	//	return
	//}
	//json.NewEncoder(w).Encode(resp)

}
