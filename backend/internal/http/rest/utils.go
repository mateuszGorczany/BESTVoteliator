package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

func respondWithError(w http.ResponseWriter, code int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	errs := strings.Split(err.Error(), "\n")
	errorResponse := map[string]interface{}{
		"code":   code,
		"errors": errs,
	}
	respondWithJSON(w, code, errorResponse)
}

func respondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
