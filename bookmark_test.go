package hatena

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/garyburd/go-oauth/oauth"
)

func TestAuthenticator_GetBookmark(t *testing.T) {
	type fields struct {
		client      oauth.Client
		redirectUri string
		scopes      url.Values
		cred        *oauth.Credentials
	}
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Bookmark
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authenticator{
				client:      tt.fields.client,
				redirectUri: tt.fields.redirectUri,
				scopes:      tt.fields.scopes,
				cred:        tt.fields.cred,
			}
			got, err := a.GetBookmark(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticator.GetBookmark() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticator.GetBookmark() = %v, want %v", got, tt.want)
			}
		})
	}
}

var bookmarkEntry = `
{
count: 1,
favicon_url: "https://cdn-ak.favicon.st-hatena.com/?url=https%3A%2F%2Fapi.nasa.gov%2F",
eid: 273202811,
image_last_editor: "",
entry_url: "http://b.hatena.ne.jp/entry/s/api.nasa.gov/",
root_url: "https://api.nasa.gov/",
is_invalid_url: false,
smartphone_app_entry_url: "http://b.hatena.ne.jp/bookmarklet.touch?mode=comment&iphone_app=1&url=https%3A%2F%2Fapi.nasa.gov%2F",
title_last_editor: "",
image_url: "",
url: "https://api.nasa.gov/",
title: "NASA Open APIs",
has_asin: false,
image_hatena_url: ""
}`
