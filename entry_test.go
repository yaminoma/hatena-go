package hatena

import (
	"reflect"
	"testing"
)

func TestEntryInfo(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *EntryInformation
		wantErr bool
	}{
		{
			name: "TestClient_EntryInfo",
			args: args{
				urlStr: "http://hatena.ne.jp",
			},
			want: &EntryInformation{
				Eid:        44054,
				Title:      "はてな",
				Count:      55,
				Url:        "http://hatena.ne.jp/",
				EntryUrl:   "http://b.hatena.ne.jp/entry/hatena.ne.jp/",
				Screenshot: "http://screenshot.hatena.ne.jp/images/200x150.jpg",
				Bookmarks: []BookmarkUser{
					{
						User:      "test-user1",
						Comment:   "test-comment1",
						Timestamp: "2006/02/20 18:12:52",
						Tags:      []string{"hatena", "はてな"},
					},
					{
						User:      "test-user2",
						Comment:   "test-comment2",
						Timestamp: "2005/03/26 23:37:16",
						Tags:      []string{"hatena", "はてな"},
					},
				},
				RelatedEntries: []RelatedEntry{
					{
						Eid:      7405822,
						Title:    "hatena-related-website",
						Count:    1,
						Url:      "http://hatena-related-website.com",
						EntryUrl: "http://b.hatena.ne.jp/entry/hatena-related-website",
					},
					{
						Eid:      40214150,
						Title:    "Delicious Transition",
						Count:    6,
						Url:      "http://www.delicious.com/help/transition",
						EntryUrl: "http://b.hatena.ne.jp/entry/www.delicious.com/help/transition",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientString(200, entryJson)
			got, err := c.EntryInfo(tt.args.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntryInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntryInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

var entryJson = `
{
	"related": [{
		"count": 1,
		"url": "http://hatena-related-website.com",
		"eid": 7405822,
		"title": "hatena-related-website",
		"entry_url": "http://b.hatena.ne.jp/entry/hatena-related-website"
	}, {
		"count": 6,
		"url": "http://www.delicious.com/help/transition",
		"eid": 40214150,
		"title": "Delicious Transition",
		"entry_url": "http://b.hatena.ne.jp/entry/www.delicious.com/help/transition"
	}],
	"count": 55,
	"bookmarks": [{
		"comment": "test-comment1",
		"timestamp": "2006/02/20 18:12:52",
		"user": "test-user1",
		"tags": ["hatena", "はてな"]
	}, {
		"comment": "test-comment2",
		"timestamp": "2005/03/26 23:37:16",
		"user": "test-user2",
		"tags": ["hatena", "はてな"]
	}],
	"url": "http://hatena.ne.jp/",
	"eid": 44054,
	"title": "はてな",
	"screenshot": "http://screenshot.hatena.ne.jp/images/200x150.jpg",
	"entry_url": "http://b.hatena.ne.jp/entry/hatena.ne.jp/"
}
`
