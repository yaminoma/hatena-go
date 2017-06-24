package hatena

import (
	"net/http"
	"testing"
)

func TestClient_Count(t *testing.T) {
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
		want    int
		wantErr bool
	}{
		0: {
			name: "TestClient_Count",
			fields: fields{
				http: new(http.Client),
			},
			args: args{
				urlStr: "http://www.hatena.ne.jp/",
			},
			want:    5818,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				http: tt.fields.http,
			}
			got, err := c.Count(tt.args.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got < tt.want {
				t.Errorf("Client.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
