package main

import (
	"log"
	"os"
	"flag"
	"net/http"

	"github.com/dev-urandom/graft"
	"github.com/bmizerany/pat"
)

var (
	logger *log.Logger
	port string
	host string
	server *graft.Server
	serverStarted bool
)

func init() {
	flag.StringVar(&port, "port", "7777", "port")
	flag.StringVar(&host, "host", "localhost", "host")
	flag.Parse()

	server = graft.New()
	serverStarted = false
	logger = log.New(os.Stdin, "["+host+":"+port+"] ", log.LstdFlags)
}

func main() {
	http.Handle("/", routes(server))

	logger.Println("Listening...")

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes(server *graft.Server) http.Handler {
	handle := pat.New()

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

		if !serverStarted {
			server.Start()
			serverStarted = true

			res.respond(200, JsonData{"message": "server starting"})
		} else {
			res.respond(400, JsonData{"message": "server already started"})
		}
	}))

	return handle
}
