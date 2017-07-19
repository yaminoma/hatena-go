package hatena

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/garyburd/go-oauth/oauth"
)

func TestAuthenticator_GetProfile(t *testing.T) {
	type fields struct {
		client      oauth.Client
		redirectUri string
		scopes      url.Values
		cred        *oauth.Credentials
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Profile
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
			got, err := a.GetProfile()
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticator.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticator.GetProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
