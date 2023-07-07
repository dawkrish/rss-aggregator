package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/krishnanshagarwal112/rss-aggregator/internal/database"
)

func (cfg *apiConfig)feedPost(w http.ResponseWriter, r *http.Request, user database.User) {
	type request struct{
		Name string `json:"name"`
		Url string `json:"url"`
	}
	requestBody := request{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	feed,err := cfg.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: requestBody.Name,
		Url: requestBody.Url,
		UserID: user.ID,
	})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	RespondWithJson(w,200,databaseFeedToFeed(feed))
}

func (cfg *apiConfig) feedGet(w http.ResponseWriter, r* http.Request){
	feeds,err:=cfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w,http.StatusInternalServerError,"can't fetch feeds")
		return
	}

	RespondWithJson(w,200,feeds)
}