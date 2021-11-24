package main

import (
	"net/http"
	"encoding/json"
)

type Coaster struct {
	Name string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID string `json:"id"`
	InPark string `json:"inPark"`
	Height int `json:"height"`
}

type coasterHandlers struct {
	store map[string]Coaster
}

func main() {
	h := NewCoasterHandlers()
	http.HandleFunc("/coasters", h.coastersHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func (h *coasterHandlers) coastersHandler(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))
	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func NewCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1": Coaster{
				Name: "abc",
				Manufacturer: "abc",
				ID: "id1",
				InPark: "abc",
				Height: 123,
			},
		},
	}
}
