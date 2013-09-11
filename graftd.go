package main

import (
	"io"
	"log"
	"net/http"

	"github.com/dev-urandom/graft"
	"github.com/bmizerany/pat"
)


func main() {
	host := "localhost"
	port := "7777"
	server := graft.New()

	http.Handle("/", routes())

	log.Println("Server State:", server.State)
	log.Println("Listening on localhost:7777")

	err := http.ListenAndServe(host+":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes() http.Handler {
	handle := pat.New()

	handle.Get("/status", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, "OK")
	}))

	return handle
}
