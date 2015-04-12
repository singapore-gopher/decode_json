package main

import (
	`flag`
	`github.com/gorilla/mux`
	"github.com/singapore-gophers/decode_json/stream"
	`log`
	`net/http`
)

func main() {
	flag.Parse()
	setupChallenge()

	go saveChallenge()
	go serveStream()
	http.Handle(`/`, getRouter())
	err := http.ListenAndServe(`:4000`, nil)
	if err != nil {
		log.Fatalf(`Could not start web server; err=%v`, err)
	}
}

func serveStream() {
	stream.Port = "4001"
	status := make(chan int)
	go stream.Serve(status)
	select {
	case s := <-status:
		if s != 0 {
			log.Println("Failed to start stream server")
			return
		}
	}

	log.Printf("stream server listening on port %s", stream.Port)

	for {
		select {
		case s := <-status:
			if s == 2 {
				// this is not implemented, so (for now) it waits forever.
				return
			}
		}
	}
}

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(`/register.json`, register).Methods(`POST`)
	router.HandleFunc(`/leaderboard.json`, leaderboard).Methods(`GET`)

	stage1 := router.PathPrefix(`/stage1`).Subrouter()
	stage1.HandleFunc(`/data.json`, dataHandler(Stage1)).Methods(`GET`)
	stage1.HandleFunc(`/submit.json`, stage1Handler).Methods(`POST`)

	stage2 := router.PathPrefix(`/stage2`).Subrouter()
	stage2.HandleFunc(`/data.json`, placeholder).Methods(`GET`)
	stage2.HandleFunc(`/submit.json`, placeholder).Methods(`POST`)

	stage3 := router.PathPrefix(`/stage3`).Subrouter()
	stage3.HandleFunc(`/data.json`, placeholder).Methods(`GET`)
	stage3.HandleFunc(`/submit.json`, placeholder).Methods(`POST`)

	return router
}

func placeholder(resp http.ResponseWriter, req *http.Request) {
}
