package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

func IsEnvProd() bool {
	return os.Getenv("ENV") == "production"
}

func AppEnv() string {
	return os.Getenv("ENV")
}

func JSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
