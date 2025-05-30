package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func makeRouter(store *QuoteStore) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var q Quote
			if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
				http.Error(w, "invalid input", http.StatusBadRequest)
				return
			}
			q = store.Add(q)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(q)
		} else {
			author := r.URL.Query().Get("author")
			var quotes []Quote
			if author != "" {
				quotes = store.FilterByAuthor(author)
			} else {
				quotes = store.GetAll()
			}
			if quotes == nil {
				quotes = []Quote{}
			}
			json.NewEncoder(w).Encode(quotes)
		}
	}).Methods("GET", "POST")

	r.HandleFunc("/quotes/random", func(w http.ResponseWriter, r *http.Request) {
		q, ok := store.GetRandom()
		if !ok {
			http.Error(w, "no quotes available", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(q)
	}).Methods("GET")

	r.HandleFunc("/quotes/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if ok := store.DeleteByID(id); !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}).Methods("DELETE")

	return r
}
