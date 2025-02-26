package notes

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
	sr := router.PathPrefix("/notes").Subrouter()
	sr.Use(middleware.NoteValidatorMiddleware)
	sr.HandleFunc("", s.ListNotes).Methods("GET")
	sr.HandleFunc("", s.CreateNote).Methods("POST")
	sr.HandleFunc("/{id}", s.GetNote).Methods("GET")
	sr.HandleFunc("/{id}", s.UpdateNote).Methods("PUT")
	sr.HandleFunc("/{id}", s.DeleteNote).Methods("DELETE")
}

func (s *Service) ListNotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notes, err := s.queries.ListNotes(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notes)
}

func (s *Service) CreateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params models.CreateNoteParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	note, err := s.queries.CreateNote(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(note)
}

func (s *Service) GetNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	note, err := s.queries.NoteById(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(note)
}

func (s *Service) UpdateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	var params models.UpdateNoteParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params.ID = int32(id)
	note, err := s.queries.UpdateNote(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(note)
}

func (s *Service) DeleteNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	err := s.queries.DeleteNote(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
