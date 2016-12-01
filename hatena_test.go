package hatena_go

import (
	"net/http"
	"os"
	"testing"
)

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

// Returns a client whose requests will always return
// a response with the specified status code and a body
// that is read from the specified file.
func testClientFile(code int, filename string) *Client {
	return &Client{
		http: &http.Client{
			Transport: newFileRoundTripper(code, filename),
		},
	}
}
