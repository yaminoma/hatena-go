package hatena

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuthenticator_GetProfile(t *testing.T) {
	tests := []struct {
		name    string
		want    *Profile
		wantErr bool
	}{
		{
			name: "TestAuthenticator_GetProfile",
			want: &Profile{
				IsOauthMixiCheck:    false,
				MixiCheckChecked:    "inherit",
				IgnoresRegex:        "",
				IsOauthEvernote:     true,
				TwitterChecked:      "inherit",
				Plususer:            false,
				BookmarkCount:       1657,
				UserPageVersion:     "1",
				DefaultSharedLinkTo: "origin",
				EvernoteChecked:     "inherit",
				FacebookChecked:     "inherit",
				Name:                "test-user1",
				Private:             false,
				Rkm:                 "83EJ/kIpQzx1ilOaMVIgFg",
				IsOauthTwitter:      true,
				Login:               true,
				Rks:                 "c69d6370fb4f5391ac4470be29c66060f38feb12",
				IsOauthFacebook:     true,
				IsStaff:             false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := testAuthClientFile(http.StatusOK, "test_data/profile.txt")
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
