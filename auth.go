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
	client      oauth.Client
	httpClient  *http.Client
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
		TemporaryCredentialRequestURI: "https://www.hatena.ne.jp/oauth/initiate",
		ResourceOwnerAuthorizationURI: "https://www.hatena.ne.jp/oauth/authorize",
		TokenRequestURI:               "https://www.hatena.ne.jp/oauth/token",
	}

	scopeParam := url.Values{}
	for _, v := range scopes {
		scopeParam.Add("scope", v)
	}

	return &Authenticator{
		client:      oauthClient,
		redirectUri: redirectUri,
		scopes:      scopeParam,
	}
}

func (a *Authenticator) AuthURL(w http.ResponseWriter, r *http.Request) (string, error) {
	tempCred, err := a.client.RequestTemporaryCredentials(nil, a.redirectUri, a.scopes)
	if err != nil {
		return "", err
	}

	//Save Session
	s := session.Get(r)
	s[tempCredKey] = tempCred
	if err := session.Save(w, r, s); err != nil {
		return "", err
	}

	return a.client.AuthorizationURL(tempCred, nil), nil
}

func (a *Authenticator) Token(w http.ResponseWriter, r *http.Request) (*oauth.Credentials, error) {

	s := session.Get(r)
	tempCred, ok := s[tempCredKey].(*oauth.Credentials)
	if !ok {
		return nil, errors.New("saved session type is invalid")
	}
	if tempCred == nil || tempCred.Token != r.FormValue("oauth_token") {
		http.Error(w, "Unknown oauth_token.", 500)
	}

	tokenCred, _, err := a.client.RequestToken(nil, tempCred, r.FormValue("oauth_verifier"))
	if err != nil {
		log.Fatal(w, "Error getting request token, "+err.Error(), 500)
	}
	a.cred = tokenCred

	delete(s, tempCredKey)
	s[tokenCredKey] = tokenCred
	if err := session.Save(w, r, s); err != nil {
		return nil, err
	}

	return a.cred, nil
}

// apiGet issues a GET request to the Hatena API and decodes the response JSON to data.
func (a *Authenticator) apiGet(urlStr string, form url.Values, result interface{}) error {
	resp, err := a.client.Get(a.httpClient, a.cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, result)
}

// apiPost issues a POST request to the Hatena API and decodes the response JSON to data.
func (a *Authenticator) apiPost(urlStr string, form url.Values, result interface{}) error {
	resp, err := a.client.Post(a.httpClient, a.cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, result)
}

// apiDelete issues a DELETE request to the Hatena API and decodes the response JSON to data.
func (a *Authenticator) apiDelete(urlStr string, form url.Values, result interface{}) error {
	resp, err := a.client.Delete(a.httpClient, a.cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resultByte, _ := ioutil.ReadAll(resp.Body)

	//削除成功の場合は204(http.StatusNoContent)が返却される
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, resultByte)
	}
	return json.NewDecoder(strings.NewReader(string(resultByte))).Decode(result)
}

// decodeResponse decodes the JSON response from the Hatena API.
func decodeResponse(resp *http.Response, result interface{}) error {
	resultByte, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, resultByte)
	}
	fmt.Println(string(resultByte))
	return json.NewDecoder(strings.NewReader(string(resultByte))).Decode(result)
}
