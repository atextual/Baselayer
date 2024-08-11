package main

import (
	"BaseLayer/handler"
	"BaseLayer/middleware"
	"BaseLayer/model"
	"BaseLayer/repo"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	var projects model.Projects = map[string]*model.Project{}

	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalln("Failed to open config.yml")
	}

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatalln("Failed to read content from config.yml")
	}

	if err := yaml.Unmarshal(data, &projects); err != nil {
		log.Fatalln("Failed to parse content from config.yml")
	}

	err = f.Close()
	if err != nil {
		log.Fatalln("Failed to close config.yml, fatal IO error: " + err.Error())
	}

	log.Println("Running internal BaseLayer database validation")
	initialiseDatabase()
	cxn, err := repo.GetConnection()
	db := cxn.Db

	log.Println("Located " + strconv.Itoa(len(projects)) + " projects in config")
	for key, project := range projects {
		project.Name = key // Slightly hacky workaround to get the key from the map assigned to the struct
		log.Println("Project name: " + project.Name + " (" + path.Join(project.ProjectDirectory, project.SqlDirectory) + ")")

		normalisedProjectName := strings.ToUpper(key) + "_DATABASE"
		row := db.QueryRow("SELECT COUNT(*) FROM databases WHERE normalised_name = ?", normalisedProjectName)

		if err != nil {
			log.Println(err.Error())
		}

		var recordCount int
		err = row.Scan(&recordCount)

		if err != nil {
			log.Println(err.Error())
		}

		if recordCount == 1 {
			log.Println("Project " + project.Name + " already in db, skipping")
		} else if recordCount == 0 {
			log.Println("Project " + project.Name + " currently untracked, adding")

			// Some columns will not be provided in config (e.g. normalised columns) and some are optional
			// (e.g. database name), make a best guess at proving some default values before proceeding
			project.Database.Name = project.Name + "_database"
			project.Database.NormalisedName = strings.ToUpper(project.Name + "_database")
			project.Database.NormalisedDriver = strings.ToUpper(project.Database.Driver)

			_, err := repo.AddDatabase(db, project)
			if err != nil {
				log.Println("Failed to add project " + project.Name + " to internal BaseLayer database, skipping")
				log.Println(err.Error())
			} else {
				log.Println("Successfully added project " + project.Name + " to internal BaseLayer database")
			}
		}
	}

	// @todo - dial into each database, update state in migration manager sqlite db
	// @todo - compute list of complete vs pending migrations, add to cache
	// @todo - start http listener
	r := mux.NewRouter()
	r.Use(middleware.JsonContentTypeMiddleware)

	r.HandleFunc("/", handler.RootHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/databases", handler.ListDatabases).Methods("GET")
	r.HandleFunc("/databases/{id:[0-9]+}", handler.GetDatabase).Methods("GET")
	r.HandleFunc("/databases", handler.CreateDatabase).Methods("POST")

	log.Println("Listening on port 8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func initialiseDatabase() {
	cxn, err := repo.GetConnection()
	if err != nil {
		log.Fatalln("Failed to connect to internal BaseLayer database")
	} else {
		log.Println("Internal BaseLayer database validation successful")
	}

	db := cxn.Db
	stmt := `
		CREATE TABLE IF NOT EXISTS databases (
			id INTEGER NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			normalised_name TEXT NOT NULL,
			driver VARCHAR(11) NOT NULL,
			normalised_driver VARCHAR(11) NOT NULL,
		    username TEXT NOT NULL,
		    password TEXT NOT NULL,
		    database TEXT NOT NULL,
		    port INTEGER NOT NULL
		);
	`

	_, initErr := db.Exec(stmt)
	if initErr != nil {
		log.Fatalln("Failed to execute schema validation query: " + err.Error())
	}
}
