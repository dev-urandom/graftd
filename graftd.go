package main

import (
	"io"
	"log"
	"os"
	"flag"
	"net/http"

	"github.com/dev-urandom/graft"
	"github.com/bmizerany/pat"
)

var logger *log.Logger

func main() {
	port := "7777"
	flag.StringVar(&port, "port", "3333", "port")
	host := "localhost"
	flag.StringVar(&host, "host", "localhost", "host")

	flag.Parse()

	server := graft.New()
	logger = log.New(os.Stdin, "["+host+":"+port+"] ", log.LstdFlags)

	http.Handle("/", routes(server))

	logger.Println("Listening...")

	server.Start()

	err := http.ListenAndServe(host+":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes(server *graft.Server) http.Handler {
	handle := pat.New()
	started := false

	handle.Get("/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := Response{w, r}

		res.respond(200, JsonData{
			"id": server.Id,
			"term": server.Term,
			"votedFor": server.VotedFor,
			"state": server.State,
		})
	}))

	handle.Post("/start", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := Response{w, r}

		if !started {
			server.Start()
			started = true

			res.respond(200, JsonData{"message": "server starting"})
		} else {
			res.respond(400, JsonData{"message": "server already started"})
		}
	}))

	return handle
}

type Response struct {
	w http.ResponseWriter
	r *http.Request
}

func (res Response) respond(status int, data JsonData) {
	res.w.WriteHeader(status)
	io.WriteString(res.w, data.Encode())
	logger.Println(res.r.Method, res.r.URL.Path, status, data.Encode())
}
