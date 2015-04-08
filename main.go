package main

import (
	`flag`
	`github.com/gorilla/mux`
	`log`
	`net/http`
)

func main() {
	flag.Parse()
	setupChallenge()

	go saveChallenge()
	http.Handle(`/`, getRouter())
	err := http.ListenAndServe(`:4000`, nil)
	if err != nil {
		log.Fatalf(`Could not start web server; err=%v`, err)
	}
}

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(`/register.json`, register).Methods(`POST`)
	router.HandleFunc(`/leaderboard.json`, leaderboard).Methods(`GET`)

	stage1 := router.PathPrefix(`/stage1`).Subrouter()
	stage1.HandleFunc(`/data.json`, placeholder).Methods(`GET`)
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
