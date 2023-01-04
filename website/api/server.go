package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router
	userList []User
}

func NewServer() *Server {
	s := &Server{
		Router:   mux.NewRouter(),
		userList: []User{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/userList", s.listUserItem()).Methods("GET")
	s.HandleFunc("/userList", s.createUserItem()).Methods("POST")
	s.HandleFunc("/userList", s.removeUserItem()).Methods("DELETE")
}

func (s *Server) createUserItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		u.ID = uuid.New()
		s.userList = append(s.userList, u)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listUserItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.userList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeUserItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for u, user := range s.userList {
			if user.ID == id {
				s.userList = append(s.userList[:u], s.userList[u+1:]...)
				break
			}
		}
	}
}
