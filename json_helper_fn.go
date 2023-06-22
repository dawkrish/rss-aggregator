package main

import (
	"encoding/json"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	jsonResponse, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	http.Error(w,msg,code)
}
