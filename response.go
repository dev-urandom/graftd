package main

import (
	"io"
	"net/http"
)

type Response struct {
	w http.ResponseWriter
	r *http.Request
}

func (res Response) respond(status int, data JsonData) {
	res.w.WriteHeader(status)
	io.WriteString(res.w, data.Encode())
	logger.Println(res.r.Method, res.r.URL.Path, status, data.Encode())
}
