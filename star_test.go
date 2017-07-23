package hatena

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_UserStar(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserStars
		wantErr bool
	}{
		{
			name: "TestClient_UserStar",
			args: args{
				username: "jkondo",
			},
			want: &UserStars{
				Count: struct {
					Green  int    `json:"green"`
					Blue   int    `json:"blue"`
					Red    int    `json:"red"`
					Purple int    `json:"purple"`
					Yellow string `json:"yellow"`
				}{
					Green:  284,
					Blue:   30,
					Red:    104,
					Purple: 4,
					Yellow: "10171",
				},
				Title:     "はてなブックマーク - jkondoのブックマーク",
				URI:       "http://b.hatena.ne.jp/jkondo/",
				StarCount: 10593,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/star_user.txt")
			got, err := c.UserStar(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UserStar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.UserStar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetStar(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *EntryStars
		wantErr bool
	}{
		{
			name: "TestClient_GetStar",
			args: args{
				urlStr: "http://b.hatena.ne.jp/jkondo/20160330#bookmark-283654293",
			},
			want: &EntryStars{
				Entries: []EntryPost{
					{
						Stars: []Star{
							{
								Count: 393,
								Quote: "comment1",
								Name:  "test-user1",
							},
							{
								Count: 1,
								Quote: "comment2",
								Name:  "test-user2",
							},
						},
						ColoredStars: []ColoredStars{
							{
								Stars: []Star{
									{
										Quote: "comment3",
										Name:  "test-user3",
										Count: 1,
									},
									{
										Quote: "comment4",
										Name:  "test-user4",
										Count: 1,
									},
								},
								Color: "green",
							},
						},
						CanComment: 0,
						URI:        "http://jkondo.hatenablog.com/entry/20060422/1145674096",
					},
				},
				CanComment: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/star.txt")
			got, err := c.GetStar(tt.args.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetStar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetStar() = %v, want %v", got, tt.want)
			}
		})
	}
}
