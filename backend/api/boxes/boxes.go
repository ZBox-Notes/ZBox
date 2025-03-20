package boxes

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ZBox-Notes/ZBox/backend/middleware"
	models "github.com/ZBox-Notes/ZBox/backend/models/generated_model"

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
	slog.Info("Registering boxes handler")
	sr := router.PathPrefix("/boxes").Subrouter()
	sr.Use(middleware.BoxValidatorMiddleware)
	sr.HandleFunc("", s.ListBoxes).Methods("GET")
	sr.HandleFunc("", s.CreateBox).Methods("POST")
	sr.HandleFunc("/{id}", s.GetBox).Methods("GET")
	sr.HandleFunc("/{id}", s.UpdateBox).Methods("PUT")
	sr.HandleFunc("/{id}", s.DeleteBox).Methods("DELETE")
}

func (s *Service) ListBoxes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	boxes, err := s.queries.ListBoxes(ctx)
	if err != nil {
		slog.Error("Error listing boxes", "Error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(boxes)
}

func (s *Service) CreateBox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var params models.CreateBoxParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		slog.Error("Error decoding request body", "Error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	box, err := s.queries.CreateBox(ctx, params)
	if err != nil {
		slog.Error("Error creating box", "Error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(box)
}

func (s *Service) GetBox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		slog.Error("Error parsing id", "Error", parseIntErr.Error())
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	box, err := s.queries.BoxById(ctx, int32(id))
	if err != nil {
		slog.Error("Error getting box with ID", "ID", int32(id), "Error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(box)
}

func (s *Service) UpdateBox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		slog.Error("Error parsing id", "Error", parseIntErr.Error())
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	var params models.UpdateBoxParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		slog.Error("Error decoding request body", "Error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params.ID = int32(id)
	box, err := s.queries.UpdateBox(ctx, params)
	if err != nil {
		slog.Error("Error updating box with ID", "ID", int32(id), "Error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(box)
}

func (s *Service) DeleteBox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := mux.Vars(r)["id"]
	id, parseIntErr := strconv.ParseInt(idStr, 10, 32)
	if parseIntErr != nil {
		slog.Error("Error parsing id", "Error", parseIntErr.Error())
		http.Error(w, parseIntErr.Error(), http.StatusBadRequest)
		return
	}
	err := s.queries.DeleteBox(ctx, int32(id))
	if err != nil {
		slog.Error("Error deleting box with ID", "ID", int32(id), "Error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
