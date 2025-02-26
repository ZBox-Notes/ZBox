package notesboxes

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
	sr := router.PathPrefix("/notesboxes").Subrouter()
	sr.Use(middleware.NoteBoxValidatorMiddleware)
	sr.HandleFunc("/notesboxes", s.ListNotesBoxes).Methods("GET")
	sr.HandleFunc("/notesboxes", s.CreateNotesBox).Methods("POST")
	sr.HandleFunc("/notesboxes", s.DeleteNotesBox).Methods("DELETE")
	sr.HandleFunc("/notesboxes/notes/{id}", s.GetNotesBoxesByNote).Methods("GET")
	sr.HandleFunc("/notesboxes/boxes/{id}", s.GetNotesBoxesByBox).Methods("GET")
}

func (s *Service) ListNotesBoxes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notesBoxes, err := s.queries.ListNotesBoxes(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notesBoxes)
}

func (s *Service) CreateNotesBox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params models.CreateNotesBoxParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	notesBox, err := s.queries.CreateNotesBox(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notesBox)
}

func (s *Service) GetNotesBoxesByNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	notesBoxes, err := s.queries.NotesBoxesByNoteId(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notesBoxes)
}

func (s *Service) GetNotesBoxesByBox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	notesBoxes, err := s.queries.NotesBoxesByBoxId(ctx, int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notesBoxes)
}

func (s *Service) DeleteNotesBox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params models.DeleteNotesBoxParams
	decodeErr := json.NewDecoder(r.Body).Decode(&params)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}
	err := s.queries.DeleteNotesBox(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
