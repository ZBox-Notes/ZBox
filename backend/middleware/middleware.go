package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"io"
	"log/slog"
	"net/http"
	"net/mail"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintln(r.Method, r.URL))
		next.ServeHTTP(w, r)
	})
}

func UserValidatorMiddleware(next http.Handler) http.Handler {
	FULL_NAME_MAX_LENGTH, err := strconv.Atoi(os.Getenv("FULL_NAME_MAX_LENGTH"))
	if err != nil {
		slog.Error("FULL_NAME_MAX_LENGTH must be an integer")
		panic("FULL_NAME_MAX_LENGTH must be an integer")
	}
	FULL_NAME_MIN_LENGTH, err := strconv.Atoi(os.Getenv("FULL_NAME_MIN_LENGTH"))
	if err != nil {
		slog.Error("FULL_NAME_MIN_LENGTH must be an integer")
		panic("FULL_NAME_MIN_LENGTH must be an integer")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody map[string]interface{}
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Info("Failed to read request body", "Error", err.Error())
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &requestBody)
		if err == nil {
			if _, ok := requestBody["email"]; ok {
				email := requestBody["email"].(string)
				_, err := mail.ParseAddress(email)
				if err != nil {
					slog.Error("Invalid email", "Error", err.Error())
					http.Error(w, "Invalid email", http.StatusBadRequest)
					return
				}
			}

			if _, ok := requestBody["full_name"]; ok {
				fullName := requestBody["full_name"].(string)
				if len(fullName) < FULL_NAME_MIN_LENGTH {
					slog.Error(fmt.Sprintln("full_name must be at least", FULL_NAME_MIN_LENGTH, "characters long"))
					http.Error(w, fmt.Sprintln("full_name must be at least", FULL_NAME_MIN_LENGTH, "characters long"), http.StatusBadRequest)
					return
				}
				if len(fullName) > FULL_NAME_MAX_LENGTH {
					slog.Error(fmt.Sprintln("full_name must be at most", FULL_NAME_MAX_LENGTH, "characters long"))
					http.Error(w, fmt.Sprintln("full_name must be at most", FULL_NAME_MAX_LENGTH, "characters long"), http.StatusBadRequest)
					return
				}
			}
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		next.ServeHTTP(w, r)
	})
}

func NoteValidatorMiddleware(next http.Handler) http.Handler {
	NOTE_TITLE_MAX_LENGTH, err := strconv.Atoi(os.Getenv("NOTE_TITLE_MAX_LENGTH"))
	if err != nil {
		slog.Error("NOTE_TITLE_MAX_LENGTH must be an integer")
		panic("NOTE_TITLE_MAX_LENGTH must be an integer")
	}
	NOTE_CONTENT_MAX_LENGTH, err := strconv.Atoi(os.Getenv("NOTE_CONTENT_MAX_LENGTH"))
	if err != nil {
		slog.Error("NOTE_CONTENT_MAX_LENGTH must be an integer")
		panic("NOTE_CONTENT_MAX_LENGTH must be an integer")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody map[string]interface{}
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Failed to read request body", "Error", err.Error())
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &requestBody)
		if err == nil {
			if titleInter, ok := requestBody["title"]; ok {
				title := titleInter.(string)
				if len(title) > NOTE_TITLE_MAX_LENGTH {
					slog.Error(fmt.Sprintln("title must be at most", NOTE_TITLE_MAX_LENGTH, "characters long"))
					http.Error(w, fmt.Sprintln("title must be at most", NOTE_TITLE_MAX_LENGTH, "characters long"), http.StatusBadRequest)
					return
				}
			}

			if contentInter, ok := requestBody["content"]; ok {
				content := contentInter.(string)
				if len(content) < NOTE_CONTENT_MAX_LENGTH {
					slog.Error(fmt.Sprintln("content must be at most", NOTE_CONTENT_MAX_LENGTH, "characters long"))
					http.Error(w, fmt.Sprintln("content must be at most", NOTE_CONTENT_MAX_LENGTH, "characters long"), http.StatusBadRequest)
					return
				}
			}
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		next.ServeHTTP(w, r)
	})
}

func BoxValidatorMiddleware(next http.Handler) http.Handler {
	BOX_NAME_MAX_LENGTH, err := strconv.Atoi(os.Getenv("BOX_NAME_MAX_LENGTH"))
	if err != nil {
		slog.Error("BOX_NAME_MAX_LENGTH must be an integer")
		panic("BOX_NAME_MAX_LENGTH must be an integer")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody map[string]interface{}
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Failed to read request body", "Error", err.Error())
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &requestBody)
		if err == nil {
			if nameInter, ok := requestBody["name"]; ok {
				name := nameInter.(string)
				if len(name) > BOX_NAME_MAX_LENGTH {
					slog.Error(fmt.Sprintln("name must be at most", BOX_NAME_MAX_LENGTH, "characters long"))
					http.Error(w, fmt.Sprintln("name must be at most", BOX_NAME_MAX_LENGTH, "characters long"), http.StatusBadRequest)
					return
				}
			}
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		next.ServeHTTP(w, r)
	})
}

func NoteBoxValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate note box
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement authentication
		next.ServeHTTP(w, r)
	})
}
