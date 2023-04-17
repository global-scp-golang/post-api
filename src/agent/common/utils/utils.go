package utils

import (
	"encoding/json"
	"net/http"
	"post-api/src/agent/common"
)

func ResourceWithJson(w http.ResponseWriter, code int, payload interface{}) {
	responseJson, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(responseJson)
}

func ResourceWithError(w http.ResponseWriter, code int, message string) {
	// {"error": message}

	ResourceWithJson(w, code, common.CustomError{Error: message})
}
