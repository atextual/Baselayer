package handlers

import (
	"BaseLayer/models"
	"BaseLayer/repo"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func DatabaseListHandler(w http.ResponseWriter, r *http.Request) {
	var db *sqlx.DB
	var databases []models.Database
	cxn, err := repo.GetConnection()
	if err != nil {
		log.Fatalln("Failed to establish connection to internal BaseLayer database: " + err.Error())
	} else {
		db = cxn.Db
	}

	rows, err := db.Queryx("SELECT * FROM databases ORDER BY id DESC")
	if err != nil {
		log.Fatalln("Failed to retrieve database list: " + err.Error())
	}

	for rows.Next() {
		var database models.Database
		err = rows.StructScan(&database)

		if err != nil {
			log.Println(err.Error())
		}

		databases = append(databases, database)
	}

	envelope := models.ResponseEnvelope{
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
