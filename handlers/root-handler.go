package handlers

import (
	"BaseLayer/models"
	"encoding/json"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := models.ResponseEnvelope{Message: "System is running", Data: nil}

	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		_, err := w.Write([]byte(jsonError.Error()))
		if err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(jsonResponse)

	if err != nil {
		panic(err)
	}
}
