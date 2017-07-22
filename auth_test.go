package hatena

import (
	"net/http"
	_ "testing"
)

func testAuthClientFile(code int, filename string) *Authenticator {
	return &Authenticator{
		httpClient: &http.Client{
			Transport: newFileRoundTripper(code, filename),
		},
	}
}
