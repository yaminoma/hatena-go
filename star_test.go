package hatena

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_UserStar(t *testing.T) {
	type fields struct {
		http *http.Client
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UserStars
		wantErr bool
	}{
		0: {
			name: "TestClient_UserStar",
			fields: fields{
				http: new(http.Client),
			},
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
			c := &Client{
				http: tt.fields.http,
			}
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
		want    *Star
		wantErr bool
	}{
	//		0: {
	//			name: "TestClient_Star",
	//			fields: fields{
	//				http: new(http.Client),
	//			},
	//			args: args{
	//				bookmarkCommentUrl: "http://b.hatena.ne.jp/jkondo/20160330#bookmark-283654293",
	//			},
	//			want: &Star{
	//				Entries: []struct {
	//					Stars      []struct{} `json:"stars"`
	//					CanComment int        `json:"can_comment"`
	//					URI        string     `json:"uri"`
	//				}{
	//					0: {
	//						Stars: []struct {
	//							Quote string `json:"quote"`
	//							Name  string `json:"name"`
	//						}{
	//							0: {
	//								Quote: "",
	//								Name:  "kantei3",
	//							},
	//						},
	//						CanComment: 0,
	//						URI:        "http://b.hatena.ne.jp/jkondo/20160330#bookmark-283654293",
	//					},
	//				},
	//				CanComment: 0,
	//			},
	//		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				http: tt.fields.http,
			}
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
