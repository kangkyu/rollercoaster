package main

import (
	"net/http"
	"encoding/json"
	"sync"
	"io"
)

type Coaster struct {
	Name string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID string `json:"id"`
	InPark string `json:"inPark"`
	Height int `json:"height"`
}

type coasterHandlers struct {
	sync.Mutex
	store map[string]Coaster
}

func main() {
	h := NewCoasterHandlers()
	http.HandleFunc("/coasters", h.coasters)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))
	h.Lock()
	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	h.Unlock()
	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var coaster Coaster
	err = json.Unmarshal(bodyBytes, &coaster)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	h.Lock()
	defer h.Unlock()
	h.store[coaster.ID] = coaster
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
