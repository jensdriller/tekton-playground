package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

var counter = Counter{
	Count:       1,
	LastRequest: time.Now().Unix(),
}

type Counter struct {
	Count       int    `json:"count"`
	LastRequest int64  `json:"last_request"`
	Host        string `json:"host"`
}

func getCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	counter.Count++
	counter.LastRequest = time.Now().Unix()
	counter.Host, _ = os.Hostname()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(counter); err != nil {
		panic(err)
	}
}
