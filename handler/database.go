package handler

import (
	"BaseLayer/model"
	"BaseLayer/repo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func ListDatabases(w http.ResponseWriter, r *http.Request) {
	databases, err := repo.ListDatabases(nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	RespondJson(w, model.ResponseEnvelope{
		Message: "Database list retrieved successfully",
		Data:    databases,
	}, http.StatusOK)
}

func GetDatabase(w http.ResponseWriter, r *http.Request) {
	envelope := model.ResponseEnvelope{
		Message: "",
		Data:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
	}

	database, err := repo.GetDatabase(nil, id)
	if err != nil {

		if database == nil {
			envelope.Message = "No database found matching id " + strconv.Itoa(id)
			envelope.Data = nil

			RespondJson(w, envelope, http.StatusBadRequest)
			return
		}

		envelope.Message = "An internal server error occurred"
		envelope.Data = nil
		RespondJson(w, envelope, http.StatusInternalServerError)
	}

	envelope.Message = "Database lookup successful"
	envelope.Data = database

	RespondJson(w, envelope, http.StatusOK)
}
