package hatena

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_EntryInfo(t *testing.T) {
	type fields struct {
		http *http.Client
	}
	type args struct {
		urlStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *EntryInformation
		wantErr bool
	}{
		0: {
			name: "TestClient_EntryInfo",
			fields: fields{
				http: new(http.Client),
			},
			args: args{
				urlStr: "http://8pockets.hatenablog.com/entry/2013/12/30/162516",
			},
			want: &EntryInformation{
				Eid:        197534734,
				Title:      "Instagram APIを使って、シャレ乙な写真を集める。 - 8pocket's Space",
				Count:      3,
				Url:        "http://8pockets.hatenablog.com/entry/2013/12/30/162516",
				EntryUrl:   "http://b.hatena.ne.jp/entry/8pockets.hatenablog.com/entry/2013/12/30/162516",
				Screenshot: "http://screenshot.hatena.ne.jp/images/200x150/0/e/b/e/7/cba734c2885e8e563d6189e3b49f650409c.jpg",
				Bookmarks: []BookmarkUser{
					0: {
						User:      "s-ooizumi0811",
						Comment:   "",
						Timestamp: "2015/10/22 17:08:46",
						Tags:      []string{"instagram", "API", "あとで読む"},
					},
					1: {
						User:      "ichao130",
						Comment:   "",
						Timestamp: "2014/07/11 15:38:13",
						Tags:      []string{"Instagram"},
					},
					2: {
						User:      "canamen",
						Comment:   "",
						Timestamp: "2014/05/31 12:44:12",
						Tags:      []string{"instagram"},
					},
				},
				RelatedEntries: []RelatedEntry{
					0: {
						Eid:      216875061,
						Title:    "Successfully Tracking an Instagram Campaign | TOTEMS",
						Count:    1,
						Url:      "http://analytics.totems.co/blog/tracking-instagram-campaign/",
						EntryUrl: "http://b.hatena.ne.jp/entry/analytics.totems.co/blog/tracking-instagram-campaign/",
					},
					1: {
						Eid:      250812446,
						Title:    "インスタグラムが日本での広告開始を告知 | Fashionsnap.com",
						Count:    1,
						Url:      "http://www.fashionsnap.com/the-posts/2015-05-10/instagram-add/?PageSpeed=noscript",
						EntryUrl: "http://b.hatena.ne.jp/entry/www.fashionsnap.com/the-posts/2015-05-10/instagram-add/?PageSpeed=noscript",
					},
					2: {
						Eid:      251648600,
						Title:    "インスタグラムのフォロワー1500%！飲食店の活用事例 | 100のアウトプット",
						Count:    1,
						Url:      "http://growth-ideas.com/socialmedia-marketing-3",
						EntryUrl: "http://b.hatena.ne.jp/entry/growth-ideas.com/socialmedia-marketing-3",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				http: tt.fields.http,
			}
			got, err := c.EntryInfo(tt.args.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EntryInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.EntryInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		0: {
			name: "TestClient_EntryInfo",
			args: args{
				urlStr: "http://8pockets.hatenablog.com/entry/2013/12/30/162516",
			},
			want: &EntryInformation{
				Eid:        197534734,
				Title:      "Instagram APIを使って、シャレ乙な写真を集める。 - 8pocket's Space",
				Count:      3,
				Url:        "http://8pockets.hatenablog.com/entry/2013/12/30/162516",
				EntryUrl:   "http://b.hatena.ne.jp/entry/8pockets.hatenablog.com/entry/2013/12/30/162516",
				Screenshot: "http://screenshot.hatena.ne.jp/images/200x150/0/e/b/e/7/cba734c2885e8e563d6189e3b49f650409c.jpg",
				Bookmarks: []BookmarkUser{
					0: {
						User:      "s-ooizumi0811",
						Comment:   "",
						Timestamp: "2015/10/22 17:08:46",
						Tags:      []string{"instagram", "API", "あとで読む"},
					},
					1: {
						User:      "ichao130",
						Comment:   "",
						Timestamp: "2014/07/11 15:38:13",
						Tags:      []string{"Instagram"},
					},
					2: {
						User:      "canamen",
						Comment:   "",
						Timestamp: "2014/05/31 12:44:12",
						Tags:      []string{"instagram"},
					},
				},
				RelatedEntries: []RelatedEntry{
					0: {
						Eid:      216875061,
						Title:    "Successfully Tracking an Instagram Campaign | TOTEMS",
						Count:    1,
						Url:      "http://analytics.totems.co/blog/tracking-instagram-campaign/",
						EntryUrl: "http://b.hatena.ne.jp/entry/analytics.totems.co/blog/tracking-instagram-campaign/",
					},
					1: {
						Eid:      250812446,
						Title:    "インスタグラムが日本での広告開始を告知 | Fashionsnap.com",
						Count:    1,
						Url:      "http://www.fashionsnap.com/the-posts/2015-05-10/instagram-add/?PageSpeed=noscript",
						EntryUrl: "http://b.hatena.ne.jp/entry/www.fashionsnap.com/the-posts/2015-05-10/instagram-add/?PageSpeed=noscript",
					},
					2: {
						Eid:      251648600,
						Title:    "インスタグラムのフォロワー1500%！飲食店の活用事例 | 100のアウトプット",
						Count:    1,
						Url:      "http://growth-ideas.com/socialmedia-marketing-3",
						EntryUrl: "http://b.hatena.ne.jp/entry/growth-ideas.com/socialmedia-marketing-3",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EntryInfo(tt.args.urlStr)
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
