package handlers

import "net/http"

func Broker(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Payload{
		Error:   false,
		Message: "Hit the broker",
	})
}
