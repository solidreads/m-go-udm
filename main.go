package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	userEnd := user
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/courses", getCourses).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
func getCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
