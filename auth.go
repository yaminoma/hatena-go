package hatena

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

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

var oauthClient = NewAuthenticator("", "", "", []string{})

type Authenticator struct {
	config      oauth.Client
	redirectUri string
	scopes      url.Values
	cred        *oauth.Credentials
}

func NewAuthenticator(consumerKey string, consumerSecret string, redirectUri string, scopes []string) *Authenticator {
	oauthClient := oauth.Client{
		Credentials: oauth.Credentials{
			Token:  consumerKey,
			Secret: consumerSecret,
		},
		TemporaryCredentialRequestURI: "https://www.hatena.com/oauth/initiate",
		ResourceOwnerAuthorizationURI: "https://www.hatena.com/oauth/authorize",
		TokenRequestURI:               "https://www.hatena.com/oauth/token",
	}

	scopeParam := url.Values{}
	for _, v := range scopes {
		scopeParam.Add("scope", v)
	}

	return &Authenticator{
		config:      oauthClient,
		redirectUri: redirectUri,
		scopes:      scopeParam,
	}
}

func (auth *Authenticator) AuthURL(w http.ResponseWriter, r *http.Request) (string, error) {
	tempCred, err := auth.config.RequestTemporaryCredentials(nil, auth.redirectUri, auth.scopes)
	if err != nil {
		return "", err
	}

	//Save Session
	s := session.Get(r)
	s[tempCredKey] = tempCred
	if err := session.Save(w, r, s); err != nil {
		return "", err
	}

	return auth.config.AuthorizationURL(tempCred, nil), nil
}

func (auth *Authenticator) Token(w http.ResponseWriter, r *http.Request) (*oauth.Credentials, error) {

	s := session.Get(r)
	tempCred, ok := s[tempCredKey].(*oauth.Credentials)
	if !ok {
		return nil, errors.New("saved session type is invalid")
	}
	if tempCred == nil || tempCred.Token != r.FormValue("oauth_token") {
		http.Error(w, "Unknown oauth_token.", 500)
	}

	tokenCred, _, err := auth.config.RequestToken(nil, tempCred, r.FormValue("oauth_verifier"))
	if err != nil {
		log.Fatal(w, "Error getting request token, "+err.Error(), 500)
	}
	auth.cred = tokenCred

	delete(s, tempCredKey)
	s[tokenCredKey] = tokenCred
	if err := session.Save(w, r, s); err != nil {
		return nil, err
	}

	return auth.cred, nil
}

// authHandler reads the auth cookie and invokes a handler with the result.
type AuthHandler struct {
	Handler  func(w http.ResponseWriter, r *http.Request)
	Optional bool
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler(w, r)
}

// apiGet issues a GET request to the Hatena API and decodes the response JSON to data.
func (auth *Authenticator) apiGet(urlStr string, form url.Values, result interface{}) error {
	resp, err := auth.config.Get(nil, auth.cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, result)
}

// apiPost issues a POST request to the Hatena API and decodes the response JSON to data.
func (auth *Authenticator) apiPost(urlStr string, form url.Values, result interface{}) error {
	resp, err := auth.config.Post(nil, auth.cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, result)
}

// apiDelete issues a DELETE request to the Hatena API and decodes the response JSON to data.
func (auth *Authenticator) apiDelete(urlStr string, form url.Values, result interface{}) error {
	resp, err := auth.config.Delete(nil, auth.cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resultByte, _ := ioutil.ReadAll(resp.Body)

	//削除成功の場合は204(http.StatusNoContent)が返却される
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, resultByte)
	}
	fmt.Println(string(resultByte))
	return json.NewDecoder(strings.NewReader(string(resultByte))).Decode(result)
}

// decodeResponse decodes the JSON response from the Hatena API.
func decodeResponse(resp *http.Response, result interface{}) error {
	resultByte, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, resultByte)
	}
	fmt.Println(string(resultByte))
	return json.NewDecoder(strings.NewReader(string(resultByte))).Decode(result)
}
