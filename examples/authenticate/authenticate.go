package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yaminoma/hatena-go"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Hatena's developer portal
// and enter this value.
const ConsumerKey = "Kea491YvUL365g=="
const ConsumerSecret = "H1dz3fGJXazEx8ccHm2BOF6dy6E="
const redirectURI = "http://localhost:8080/callback"

var (
	scopes = []string{hatena.ReadPrivate, hatena.WritePrivate, hatena.WritePublic}
	auth   = hatena.NewAuthenticator(ConsumerKey, ConsumerSecret, redirectURI, scopes)
)

func main() {
	// first start an HTTP server
	http.HandleFunc("/auth", authUrl)
	http.HandleFunc("/callback", token)
	http.HandleFunc("/profile", profile)

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
	// you can get *oauth.Credentials struct in first param.
	_, err := auth.Token(w, r)
	if err != nil {
		http.Error(w, "Error getting temp cred, "+err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/profile", 302)
}

func profile(w http.ResponseWriter, r *http.Request) {

	//resp, err := auth.GetBookmarkedEntry("https://api.nasa.gov/")
	//resp, err := auth.GetBookmark("https://api.nasa.gov/")
	//result, err := auth.DeleteBookmark("https://www.roomie.jp/2016/12/364004/")
	//if err != nil {
	//	http.Error(w, "ERROR: "+err.Error(), 500)
	//	return
	//}
	//fmt.Println(result)
	//json.NewEncoder(w).Encode(resp)

	//get your profile abount hatenabookmark
	profile, _ := auth.GetProfile()
	json.NewEncoder(w).Encode(profile)

}
