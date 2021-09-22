package models

import "net/http"

type Connection struct {
	Request  *http.Request
	Response *http.Response
}

type Options struct {
	File     string
	TotalReq int64
}
