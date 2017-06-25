hatena-go
===
[![GoDoc](https://godoc.org/github.com/8pockets/hatena-go?status.svg)](https://godoc.org/github.com/8pockets/hatena-go)
[![Build Status](https://travis-ci.org/8pockets/hatena-go.svg?branch=master)](https://travis-ci.org/8pockets/hatena-go)

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
