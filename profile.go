package hatena

import (
	"fmt"
	"net/url"
)

var (
	profileURL = fmt.Sprintf("http://api.b.hatena.ne.jp/%d/my", version)
)

type Profile struct {
	IsOauthMixiCheck    bool   `json:"is_oauth_mixi_check"`
	MixiCheckChecked    string `json:"mixi_check_checked"`
	IgnoresRegex        string `json:"ignores_regex"`
	IsOauthEvernote     bool   `json:"is_oauth_evernote"`
	TwitterChecked      string `json:"twitter_checked"`
	Plususer            bool   `json:"plususer"`
	BookmarkCount       int    `json:"bookmark_count"`
	UserPageVersion     string `json:"user_page_version"`
	DefaultSharedLinkTo string `json:"default_shared_link_to"`
	EvernoteChecked     string `json:"evernote_checked"`
	FacebookChecked     string `json:"facebook_checked"`
	Name                string `json:"name"`
	Private             bool   `json:"private"`
	Rkm                 string `json:"rkm"`
	IsOauthTwitter      bool   `json:"is_oauth_twitter"`
	Login               bool   `json:"login"`
	Rks                 string `json:"rks"`
	IsOauthFacebook     bool   `json:"is_oauth_facebook"`
	IsStaff             bool   `json:"is_staff"`
}

func GetProfile() (*Profile, error) {
	return oauthClient.GetProfile()
}

func (o *Authenticator) GetProfile() (*Profile, error) {

	form := url.Values{}

	p := &Profile{}
	err := o.apiGet(profileURL, form, p)

	return p, err
}
