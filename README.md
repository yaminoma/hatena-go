hatena-go
===
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/8pockets/hatena-go)
[![Build Status](https://img.shields.io/travis/8pockets/hatena-go.svg?style=flat-square)](https://travis-ci.org/8pockets/hatena-go)
[![Coverage Status](http://img.shields.io/coveralls/8pockets/hatena-go/master.svg?style=flat-square)](https://coveralls.io/github/8pockets/hatena-go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/8pockets/hatena-go?style=flat-square)](https://goreportcard.com/report/github.com/8pockets/hatena-go)
[![Sourcegraph for Repo Reference Count](https://img.shields.io/sourcegraph/rrc//github.com/8pockets/hatena-go.svg?style=flat-square)](https://sourcegraph.com/github.com/8pockets/hatena-go)
[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/8pockets/hatena-go/blob/master/LICENSE)

Overview
------------------
This is a Go wrapper for working with [Hatena](http://developer.hatena.ne.jp/ja/documents/bookmark/apis/rest) WEB API.  
It aims to support Hatena Bookmark Web API Endpoint.

Installation
------------------
`go get github.com/8pockets/hatena-go`

Authentication
------------------
Hatena uses OAuth1 for authentication and authorization.  
A part of Web API endpoints require an access token.

````Go
auth := hatena.NewAuthenticator(ConsumerKey, ConsumerSecret, redirectURI, scopes)

````

Hatena's Web API Authorization Guide:  
http://developer.hatena.ne.jp/ja/documents/auth/apis/oauth

API Examples
------------------
Examples of the API can be found in the [examples](https://github.com/8pockets/hatena-go/tree/master/examples) directory.

License
------------------
MIT
