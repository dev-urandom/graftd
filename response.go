package main

import (
	"io"
	"net/http"
)

type Response struct {
	w       http.ResponseWriter
	r       *http.Request
	logging bool
}

func NewResponse(w http.ResponseWriter, r *http.Request) Response {
	return Response{w, r, true}
}

func (res Response) respond(status int, data JsonData) {
	res.w.WriteHeader(status)
	io.WriteString(res.w, data.Encode())
	if res.logging {
		logger.Println(res.r.Method, res.r.URL.Path, status, data.Encode())
	}
}
