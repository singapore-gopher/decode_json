package main

import (
	`net/http`
)

func NewError(resp http.ResponseWriter, status int, err error) {
	resp.WriteHeader(status)
	resp.Write([]byte(`{"error":"` + err.Error() + `"}`))
}
