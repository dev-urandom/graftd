package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bmizerany/pat"
	"github.com/dev-urandom/graft"
)

var (
	logger        *log.Logger
	port          string
	host          string
	server        *graft.Server
	serverStarted bool
)

func init() {
	flag.StringVar(&port, "port", "7777", "port")
	flag.StringVar(&host, "host", "localhost", "host")
	flag.Parse()

	server = graft.New(host + ":" + port)
	serverStarted = false
	logger = log.New(os.Stdin, "["+host+":"+port+"] ", log.LstdFlags)
}

func main() {
	http.Handle("/", routes(server))
	http.Handle("/raft/", graft.PrefixedHttpHandler(server, "/raft"))

	logger.Println("Listening...")

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func routes(server *graft.Server) http.Handler {
	handle := pat.New()

	handle.Get("/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := NewResponse(w, r)

		res.logging = false

		res.respond(200, JsonData{
			"id":              server.Id,
			"term":            server.Term,
			"votedFor":        server.VotedFor,
			"state":           server.State,
			"lastCommitIndex": server.LastCommitIndex(),
			"lastLogIndex":    server.LastLogIndex(),
			"peers":           server.Peers,
		})
	}))

	handle.Get("/log", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := NewResponse(w, r)

		res.respond(200, JsonData{
			"log": server.Log,
		})
	}))

	handle.Post("/start", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := NewResponse(w, r)

		if !serverStarted {
			server.Start()
			serverStarted = true

			res.respond(200, JsonData{"message": "server starting"})
		} else {
			res.respond(400, JsonData{"message": "server already started"})
		}
	}))

	handle.Post("/start_election", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := NewResponse(w, r)
		server.StartElection()
		res.respond(200, JsonData{"message": "started election"})
	}))

	handle.Post("/append_entry", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := NewResponse(w, r)
		body, _ := ioutil.ReadAll(r.Body)
		newEntry := string(body)

		server.AppendEntries(newEntry)

		res.respond(201, JsonData{"message": "commited to log"})
	}))

	handle.Post("/add_peer", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := NewResponse(w, r)
		body, _ := ioutil.ReadAll(r.Body)
		newPeerURL := string(body)

		server.AddPeers(graft.HttpPeer{"http://" + newPeerURL + "/raft"})

		res.respond(201, JsonData{"message": "peer added"})
	}))

	return handle
}
