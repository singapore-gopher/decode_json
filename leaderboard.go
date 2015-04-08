package main

import (
	`encoding/json`
	`net/http`
)

func leaderboard(resp http.ResponseWriter, req *http.Request) {
	jsonChallenge.RLock()
	stats := jsonChallenge.Teams
	jsonChallenge.RUnlock()

	data, err := json.Marshal(stats) // should be okay
	if err != nil {
		NewError(resp, http.StatusInternalServerError, err)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(data)
}
