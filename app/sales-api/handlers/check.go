package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type check struct {
	log *log.Logger
}

func (c check) readness(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)
	c.log.Println(r, status)
}
