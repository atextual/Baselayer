package handler

import (
	"BaseLayer/model"
	"BaseLayer/repo"
	"encoding/json"
	"log"
	"net/http"
)

func DatabaseListHandler(w http.ResponseWriter, r *http.Request) {
	databases, err := repo.ListDatabases(nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	envelope := model.ResponseEnvelope{
		Message: "Database list retrieved successfully",
		Data:    databases,
	}

	jsonResponse, jsonError := json.Marshal(envelope)
	if jsonError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(jsonError.Error()))
		if err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)

	if err != nil {
		panic(err)
	}
}
