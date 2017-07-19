package hatena

import (
	"errors"
	"net/http"
	"os"
	"strings"
	_ "testing"
)

type stringRoundTripper struct {
	strings.Reader
	statusCode  int
	lastRequest *http.Request
}

func newStringRoundTripper(code int, s string) *stringRoundTripper {
	return &stringRoundTripper{*strings.NewReader(s), code, nil}
}

func (s stringRoundTripper) Close() error {
	return nil
}

type fileRoundTripper struct {
	*os.File
	statusCode  int
	lastRequest *http.Request
}

func newFileRoundTripper(code int, filename string) *fileRoundTripper {
	file, err := os.Open(filename)
	if err != nil {
		panic("Couldn't open file " + filename)
	}
	return &fileRoundTripper{file, code, nil}
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

func testClientString(code int, body string) *Client {
	return &Client{
		http: &http.Client{
			Transport: newStringRoundTripper(code, body),
		},
	}
}
func testClientFile(code int, filename string) *Client {
	return &Client{
		http: &http.Client{
			Transport: newFileRoundTripper(code, filename),
		},
	}
}
