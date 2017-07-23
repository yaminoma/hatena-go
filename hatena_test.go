package hatena

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
)

type stringRoundTripper struct {
	strings.Reader
	statusCode  int
	lastRequest *http.Request
}

type fileRoundTripper struct {
	*os.File
	statusCode  int
	lastRequest *http.Request
}

/*
* httpリクエストをRoundTripさせ、返り値を引数のstringにする。
 */
func testClientString(code int, body string) *Client {
	return &Client{
		http: &http.Client{
			Transport: newStringRoundTripper(code, body),
		},
	}
}

func newStringRoundTripper(code int, s string) *stringRoundTripper {
	return &stringRoundTripper{*strings.NewReader(s), code, nil}
}

func (s stringRoundTripper) Close() error {
	return nil
}

func (s *stringRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	s.lastRequest = req
	if req.Header == nil {
		if req.Body != nil {
			req.Body.Close()
		}
		return nil, errors.New("RoundTripper: nil request header")
	}
	return &http.Response{
		StatusCode: s.statusCode,
		Body:       s,
	}, nil
}

/*
* httpリクエストをRoundTripさせ、返り値をtest_dataの値にする。
 */
func testClientFile(code int, filename string) *Client {
	return &Client{
		http: &http.Client{
			Transport: newFileRoundTripper(code, filename),
		},
	}
}
func newFileRoundTripper(code int, filename string) *fileRoundTripper {
	file, err := os.Open(filename)
	if err != nil {
		panic("Couldn't open file " + filename)
	}
	return &fileRoundTripper{file, code, nil}
}

func (f *fileRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	f.lastRequest = req
	if req.Header == nil {
		if req.Body != nil {
			req.Body.Close()
		}
		return nil, errors.New("fileRoundTripper: nil request header")
	}
	return &http.Response{
		StatusCode: f.statusCode,
		Body:       f,
	}, nil
}

/*
* HTTP GET Request Test
 */
func TestClient_get(t *testing.T) {
	type args struct {
		url    string
		result interface{}
		format string
	}
	tests := []struct {
		name    string
		args    args
		want    *Entries
		wantErr bool
	}{
		{
			name: "TestClient_get",
			args: args{
				url:    "http://b.hatena.ne.jp/hotentry?mode=rss",
				result: new(Entries),
				format: "xml",
			},
			want: &Entries{
				Title:       "はてなブックマーク - 人気エントリー",
				Link:        "http://b.hatena.ne.jp/hotentry",
				Description: "最近の人気エントリー",
				Entry: []Entry{
					{
						Title:         "TopNews1",
						Link:          "https://example.com/news1/",
						Description:   "TopNews1 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
					{
						Title:         "TopNews2",
						Link:          "https://example.com/news2/",
						Description:   "TopNews2 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/hotentry_all.txt")
			err := c.get(tt.args.url, tt.args.result, tt.args.format)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.result, tt.want) {
				t.Errorf("Client.get() = %v, want %v", tt.args.result, tt.want)
			}
		})
	}
}

func TestClient_post(t *testing.T) {
	type args struct {
		url    string
		data   url.Values
		result interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *Bookmark
		wantErr bool
	}{
		{
			name: "TestClient_get",
			args: args{
				url:    "http://b.hatena.ne.jp/hotentry?mode=rss",
				data:   map[string][]string{"paramBody1": []string{"1"}, "paramBody2": []string{"2"}},
				result: new(Bookmark),
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
			c := testClientFile(http.StatusOK, "test_data/bookmark.txt")
			err := c.post(tt.args.url, tt.args.data, tt.args.result)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.post() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.result, tt.want) {
				t.Errorf("Client.post() = %v, want %v", tt.args.result, tt.want)
			}
		})
	}
}

func TestClient_delete(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestClient_delete",
			args: args{
				url: "http://b.hatena.ne.jp/hotentry?mode=rss",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientString(http.StatusNoContent, "")
			err := c.delete(tt.args.url)

			if (err != nil) != tt.wantErr {
				t.Errorf("Client.post() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
