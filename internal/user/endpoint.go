package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	Endpoints  struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirsName string `json:"first_name"`
		LastName string `json:"last_name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}

	ErrorRes struct {
		Err string `json:"error"`
	}
)

func MakeEndPoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndPoint(s),
		Get:    makeGetEndPoint(s),
		GetAll: makeGetAllEndPoint(s),
		Update: makeUpdateEndPoint(s),
		Delete: makeDeleteEndPoint(s),
	}
}

func makeCreateEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"invalid request format"})
			return
		}
		if req.FirsName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"first name is required"})
			return
		}
		if req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"last name is required"})
			return
		}
		user, err := s.Create(req.FirsName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{err.Error()})
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func makeGetEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetAllEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get All")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true}) // do something
	}
}

func makeUpdateEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeDeleteEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
