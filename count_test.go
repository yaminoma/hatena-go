package hatena

import (
	"net/http"
	"strconv"
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
		want    string
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
			want:    "5818",
			wantErr: false,
		},
	}
	userStarResponse := "5818"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientString(http.StatusOK, userStarResponse)
			got, err := c.Count(tt.args.urlStr)
			gotStr := strconv.Itoa(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStr != tt.want {
				t.Errorf("Client.Count() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}
