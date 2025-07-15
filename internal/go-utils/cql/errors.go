package cql

import (
	"net/http"

	"github.com/otelia/go-utils/httpw"
)

func HandleError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		if err.Error() == "not found" {
			httpw.Respond(w, r, http.StatusNotFound, httpw.Error{Code: http.StatusNotFound, Message: "Item not found"})
			return true
		}
		httpw.Respond(w, r, http.StatusInternalServerError, httpw.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return true
	}
	return false
}
