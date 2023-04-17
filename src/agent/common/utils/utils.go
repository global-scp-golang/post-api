package utils

import "net/http"

func ResourceWithJson(w http.ResponseWriter, code int, payload interface{}) {

}

func ResourceWithError(w http.ResponseWriter, code int, message string) {
	ResourceWithJson(w, code, message)
}
