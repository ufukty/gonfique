package http

import "net/http"

type Method string

const (
	Get     Method = http.MethodGet
	Head    Method = http.MethodHead
	Post    Method = http.MethodPost
	Put     Method = http.MethodPut
	Patch   Method = http.MethodPatch
	Delete  Method = http.MethodDelete
	Connect Method = http.MethodConnect
	Options Method = http.MethodOptions
	Trace   Method = http.MethodTrace
)
