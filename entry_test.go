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
		url string
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
				url: "https://github.com/",
			},
			want: &EntryInformation{
				Eid:        10975646,
				Title:      "GitHub",
				Count:      983,
				Url:        "https://github.com/",
				EntryUrl:   "http://b.hatena.ne.jp/entry/s/github.com/",
				Screenshot: "http://screenshot.hatena.ne.jp/images/200x150/f/d/e/b/0/3ba121c130cd7312d649e5f4fb308a2394c.jpg",
				//Bookmarks:      "",
				//RelatedEntries: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				http: tt.fields.http,
			}
			got, err := c.EntryInfo(tt.args.url)
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
