package repo

import (
	"BaseLayer/models"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

var lock = &sync.Mutex{}
var connection *models.Connection

func GetConnection() (*models.Connection, error) {
	if connection == nil {
		lock.Lock()
		defer lock.Unlock()

		db, err := sqlx.Open("sqlite3", "BaseLayer.sqlite3")
		if err != nil {
			log.Fatalln("Failed to connect to internal BaseLayer database")
			return nil, err
		} else {
			log.Println("Internal BaseLayer database connection established")
		}

		connection = &models.Connection{
			Db: db,
		}
	}

	return connection, nil
}
