package hatena

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/garyburd/go-oauth/examples/session"
	"github.com/garyburd/go-oauth/oauth"
)

const (
	HatenaReadPublic   = "read_public"
	HatenaReadPrivate  = "read_private"
	HatenaWritePublic  = "write_public"
	HatenaWritePrivate = "write_private"
	// Session state keys.
	tempCredKey  = "tempCred"
	tokenCredKey = "tokenCred"
)

type Authenticator struct {
	config      oauth.Client
	redirectUri string
	scopes      []string
	cred        *oauth.Credentials
}

func NewAuthenticator(consumerKey string, consumerSecret string, redirectUri string, scopes ...string) Authenticator {
	oauthClient := oauth.Client{
		Credentials: oauth.Credentials{
			Token:  consumerKey,
			Secret: consumerSecret,
		},
		TemporaryCredentialRequestURI: "https://www.hatena.com/oauth/initiate",
		ResourceOwnerAuthorizationURI: "https://www.hatena.com/oauth/authorize",
		TokenRequestURI:               "https://www.hatena.com/oauth/token",
	}

	return Authenticator{
		config:      oauthClient,
		redirectUri: redirectUri,
		scopes:      scopes,
	}
}

//func (auth *Authenticator) NewClient() Authenticator {
//	return Authenticator{
//		config:      oauthClient,
//		redirectUri: redirectUri,
//		scopes:      scopes,
//		cred:        Credentials,
//	}
//}

func (auth *Authenticator) AuthURL(w http.ResponseWriter, r *http.Request) string {
	scope := url.Values{"scope": auth.scopes}
	tempCred, err := auth.config.RequestTemporaryCredentials(nil, auth.redirectUri, scope)
	if err != nil {
		log.Fatal("RequestTemporaryCredentials:", err)
	}

	//Save Session
	s := session.Get(r)
	s[tempCredKey] = tempCred
	if err := session.Save(w, r, s); err != nil {
		log.Fatal(w, "Error saving session , "+err.Error(), 500)
	}

	return auth.config.AuthorizationURL(tempCred, nil)
}

func (auth *Authenticator) Token(w http.ResponseWriter, r *http.Request) *oauth.Credentials {

	s := session.Get(r)
	tempCred, _ := s[tempCredKey].(*oauth.Credentials)
	if tempCred == nil || tempCred.Token != r.FormValue("oauth_token") {
		http.Error(w, "Unknown oauth_token.", 500)
	}

	tokenCred, _, err := auth.config.RequestToken(nil, tempCred, r.FormValue("oauth_verifier"))
	if err != nil {
		log.Fatal(w, "Error getting request token, "+err.Error(), 500)
	}
	auth.cred.Token = tokenCred.Token
	auth.cred.Secret = tokenCred.Secret

	delete(s, tempCredKey)
	s[tokenCredKey] = tokenCred
	if err := session.Save(w, r, s); err != nil {
		log.Fatal(w, "Error saving session , "+err.Error(), 500)
	}

	return auth.cred
}

// authHandler reads the auth cookie and invokes a handler with the result.
type AuthHandler struct {
	handler  func(w http.ResponseWriter, r *http.Request)
	optional bool
}

// apiGet issues a GET request to the Hatena API and decodes the response JSON to data.
func (auth *Authenticator) ApiGet(urlStr string, form url.Values, data interface{}) error {
	resp, err := auth.config.Get(nil, auth.cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

// apiPost issues a POST request to the Hatena API and decodes the response JSON to data.
func (auth *Authenticator) ApiPost(urlStr string, form url.Values, data interface{}) error {
	resp, err := auth.config.Post(nil, auth.cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

// decodeResponse decodes the JSON response from the Hatena API.
func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		p, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, p)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
