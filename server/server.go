package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"codeberg.org/voyna/voyna/log4j"
	"codeberg.org/voyna/voyna/search"
)

func catchAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// Marshals 'i' and writes to ResponseWriter
func sendJSON(w http.ResponseWriter, i interface{}) error {
	var b []byte
	var err error
	if os.Getenv("PROD") == "true" {
		b, err = json.Marshal(i)
	} else {
		b, err = json.MarshalIndent(i, "", "\t")
	}
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
	log4j.Logger.Printf("Voyna server: received query %q", q)
	// TODO: handle empty queries
	rs := search.Search(q)
	err := sendJSON(w, rs)
	if err != nil {
		log4j.Logger.Printf("Voyna server: could not respond query %q: %v", q, err)
	}
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
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
