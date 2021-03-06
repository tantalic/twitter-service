package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
)

type apiResponse struct {
	Version string          `json:"version"`
	Count   int             `json:"count"`
	Tweets  []twitter.Tweet `json:"tweets"`
}

func StartApiServer(c config) error {
	http.HandleFunc("/", apiHandler)

	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	log.Printf("Listening on %s\n", addr)
	return nil
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	tweets := GetTweets()

	response := apiResponse{
		Version: Version,
		Count:   len(tweets),
		Tweets:  tweets,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-API-Version", Version)
	json.NewEncoder(w).Encode(response)
}
