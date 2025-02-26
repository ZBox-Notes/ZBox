package users

import (
	"backend/middleware"
	models "backend/models/generated_model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Service struct {
	queries *models.Queries
}

func NewService(queries *models.Queries) *Service {
	return &Service{
		queries: queries,
	}
}

func (s *Service) RegisterHandlers(router *mux.Router) {
	sr := router.PathPrefix("/users").Subrouter()
	sr.Use(middleware.UserValidatorMiddleware)
	sr.HandleFunc("", s.ListUsers).Methods("GET")
	sr.HandleFunc("", s.CreateUser).Methods("POST")
	sr.HandleFunc("/{id}", s.GetUser).Methods("GET")
	sr.HandleFunc("/{id}", s.UpdateUser).Methods("PUT")
	sr.HandleFunc("/{id}", s.DeleteUser).Methods("DELETE")
}

func (s *Service) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := s.queries.ListUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params models.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := s.queries.CreateUser(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (s *Service) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	user, err := s.queries.UserById(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (s *Service) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	var params models.UpdateUserParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params.ID = int32(id)
	user, err := s.queries.UpdateUser(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (s *Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	err := s.queries.DeleteUser(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
