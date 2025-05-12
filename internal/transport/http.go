package transport

import (
	"encoding/json"
	"net/http"

	"github.com/froppa/go-svc-template/internal/service"
)

// NewHTTPHandler returns an http.Handler with your routes wired up.
func NewHTTPHandler(mux *http.ServeMux, svc service.Service) http.Handler {
	// Health-check endpoint.
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	})

	// /hello endpoint with optional ?name= query parameter.
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		msg, err := svc.Hello(r.Context(), name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": msg})
	})

	return mux
}
