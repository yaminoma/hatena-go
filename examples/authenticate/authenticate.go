package main

import (
	"fmt"
	_ "log"
	"net/http"

	"github.com/8pockets/hatena-go"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Hatena's developer portal
// and enter this value.
const redirectURI = "http://localhost:8080/callback"
const ConsumerKey = "2wDjTLiLsfRplA=="
const ConsumerSecret = "ts+bQgDIp/GI1I5q+v8Ca+12pA0="

var (
	auth = hatena.NewAuthenticator(ConsumerKey, ConsumerSecret, redirectURI, hatena.HatenaReadPrivate)
)

func main() {
	// first start an HTTP server
	http.HandleFunc("/auth", authUrl)

	http.HandleFunc("/callback", token)
	http.Handle("/profile", hatena.AuthHandler{handler: profile})

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
	if err := auth.ApiGet(
		"http://n.hatena.com/applications/my.json",
		nil,
		&dms); err != nil {
		http.Error(w, "Error getting timeline, "+err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "%s", dms)
}
