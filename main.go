package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/solidreads/m-go-udm/internal/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	_ = godotenv.Load()

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	_ = db.AutoMigrate(&user.User{})
	userRepo := user.NewRepository(db)
	userSrvc := user.NewService(userRepo)
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
