package httpw

import (
	"encoding/json"
	"net/http"
)

func NewHTMLNotFoundHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Not Found!"))
	})
}

func NewJSONNotFoundHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := Error{Code: http.StatusNotFound, Message: "Not Found"}
		response, _ := json.Marshal(payload)

		w.WriteHeader(404)
		w.Write(response)
	})
}
