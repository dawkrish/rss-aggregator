package main

import "net/http"

func err500(w http.ResponseWriter, r *http.Request){
	RespondWithError(w,500,"Internal Server Error")
}