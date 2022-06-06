package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"codeberg.org/voyna/voyna/search"
)

func catchAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// Marshals 'i' and writes to ResponseWriter
func sendJSON(w http.ResponseWriter, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("sendJSON: error parsing to JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("sendJSON: couldn't write to ResponseWriter: %v", err)
	}

	return nil
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q") // search query
	// TODO: handle empty queries
	rs := search.Search(q)
	sendJSON(w, rs)
}

func Start() {
	mux := http.NewServeMux()
	m := map[string]func(http.ResponseWriter, *http.Request){
		"/":       catchAll,
		"/search": handleSearch, // clumsy name owing to conflict with package name
	}
	for k, v := range m {
		mux.Handle(k, http.HandlerFunc(v))
	}
	http.ListenAndServe(":8080", mux)
}

func main() {
	Start()
}
