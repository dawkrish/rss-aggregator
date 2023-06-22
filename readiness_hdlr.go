package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	payload := response{
		Status: "ok",
	}
	RespondWithJson(w,200,payload)
}