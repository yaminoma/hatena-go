package hatena

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuthenticator_GetBookmark(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		want    *Bookmark
		wantErr bool
	}{
		{
			name: "TestAuthenticator_GetBookmark",
			args: args{
				uri: "https://api.nasa.gov/",
			},
			want: &Bookmark{
				CommentRaw:      "[nasa][api]NASA API",
				Private:         false,
				Eid:             273202811,
				CreatedEpoch:    1460460002,
				Tags:            []string{"api", "nasa"},
				Permalink:       "http://b.hatena.ne.jp/eightpockets/20160412#bookmark-273202811",
				Comment:         "NASA API",
				CreatedDatetime: "2016-04-12T20:20:02+09:00",
				User:            "test-user1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := testAuthClientFile(http.StatusOK, "test_data/bookmark.txt")
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

func TestAuthenticator_AddBookmark(t *testing.T) {
	type args struct {
		br BookmarkForm
	}
	tests := []struct {
		name    string
		args    args
		want    *Bookmark
		wantErr bool
	}{
		{
			name: "TestAuthenticator_AddBookmark",
			args: args{
				br: BookmarkForm{
					uri:           "https://api.nasa.gov/",
					comment:       "[nasa][api]NASA API",
					tags:          []string{"api", "nasa"},
					post_twitter:  false,
					post_facebook: false,
					post_mixi:     false,
					post_evernote: false,
					send_mail:     false,
					private:       false,
				},
			},
			want: &Bookmark{
				CommentRaw:      "[nasa][api]NASA API",
				Private:         false,
				Eid:             273202811,
				CreatedEpoch:    1460460002,
				Tags:            []string{"api", "nasa"},
				Permalink:       "http://b.hatena.ne.jp/eightpockets/20160412#bookmark-273202811",
				Comment:         "NASA API",
				CreatedDatetime: "2016-04-12T20:20:02+09:00",
				User:            "test-user1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := testAuthClientFile(http.StatusOK, "test_data/bookmark.txt")
			got, err := a.AddBookmark(tt.args.br)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticator.AddBookmark() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticator.AddBookmark() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthenticator_GetBookmarkedEntry(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		want    *BookmarkEntry
		wantErr bool
	}{
		{
			name: "TestAuthenticator_GetBookmarkedEntry",
			args: args{
				uri: "https://api.nasa.gov/",
			},
			want: &BookmarkEntry{
				Count:                 1,
				FaviconURL:            "https://cdn-ak.favicon.st-hatena.com/?url=https%3A%2F%2Fapi.nasa.gov%2F",
				Eid:                   273202811,
				ImageLastEditor:       "",
				EntryURL:              "http://b.hatena.ne.jp/entry/s/api.nasa.gov/",
				RootURL:               "https://api.nasa.gov/",
				IsInvalidURL:          false,
				SmartphoneAppEntryURL: "http://b.hatena.ne.jp/bookmarklet.touch?mode=comment&iphone_app=1&url=https%3A%2F%2Fapi.nasa.gov%2F",
				TitleLastEditor:       "",
				ImageURL:              "",
				URL:                   "https://api.nasa.gov/",
				Title:                 "NASA Open APIs",
				HasAsin:               false,
				ImageHatenaURL:        "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := testAuthClientFile(http.StatusOK, "test_data/bookmark_entry.txt")
			got, err := a.GetBookmarkedEntry(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticator.GetBookmarkedEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticator.GetBookmarkedEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
