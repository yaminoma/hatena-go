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
			want:    userStarResponseWant,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientString(http.StatusOK, userStarResponse)
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
			want:    starResponseWant,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientString(http.StatusOK, starResponse)
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

var userStarResponseWant = &UserStars{
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
}

var starResponseWant = &EntryStars{
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
}

var userStarResponse = `
{
	"count": {
		"green": 284,
		"blue": 30,
		"red": 104,
		"purple": 4,
		"yellow": "10171"
	},
	"title": "\u306f\u3066\u306a\u30d6\u30c3\u30af\u30de\u30fc\u30af - jkondo\u306e\u30d6\u30c3\u30af\u30de\u30fc\u30af",
	"uri": "http://b.hatena.ne.jp/jkondo/",
	"star_count": 10593
}`

var starResponse = `
{
	"entries": [{
		"stars": [{
			"count": 393,
			"quote": "comment1",
			"name": "test-user1"
		},{
			"count": 1,
			"quote": "comment2",
			"name": "test-user2"
		}],
		"can_comment": 0,
		"colored_stars": [{
			"stars": [{
				"count": 1,
				"quote": "comment3",
				"name": "test-user3"
			}, {
				"count": 1,
				"quote": "comment4",
				"name": "test-user4"
			}],
			"color": "green"
		}],
		"uri": "http://jkondo.hatenablog.com/entry/20060422/1145674096"
	}],
	"can_comment": 0
}`
