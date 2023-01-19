package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/solidreads/m-go-udm/internal/user"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	userSrvc := user.NewService()
	userEnd := user.MakeEndPoints(userSrvc)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users", userEnd.Update).Methods("GET")
	router.HandleFunc("/users", userEnd.Delete).Methods("PATCH")

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
