package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/krishnanshagarwal112/rss-aggregator/internal/database"
)

func userPost(w http.ResponseWriter, r *http.Request, apiCfg *apiConfig) {
	type request struct {
		Name string `json:"name"`
	}
	requestBody := request{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if len(requestBody.Name) == 0{
		RespondWithError(w,http.StatusBadRequest,"Name field is empty")
		return
	}	

	// Now we need to do things,
	// First is to create the user in database

	user,err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: requestBody.Name,
	})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	// Second is to respond with json 
	RespondWithJson(w,200,databaseUserToUser(user))
	// the func databaseUserToUser converts the user returned by the DB.CreateUser() to a struct with json tags !
}

