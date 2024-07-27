package handler

import (
	"BaseLayer/model"
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, response model.ResponseEnvelope, status int) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(jsonError.Error()))
		if err != nil {
			panic(err)
		}
	}

	w.WriteHeader(status)
	_, err := w.Write(jsonResponse)

	if err != nil {
		panic(err)
	}
}
