package main

import (
	"BaseLayer/handlers"
	"BaseLayer/middleware"
	"BaseLayer/models"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

func main() {
	var projects models.Projects = map[string]*models.Project{}

	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalln("Failed to open config.yml")
	}

	defer f.Close() // Hold off on closing the file handle until the end of main()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatalln("Failed to read content from config.yml")
	}

	if err := yaml.Unmarshal(data, &projects); err != nil {
		log.Fatalln("Failed to parse content from config.yml")
	}

	log.Println("Initialisation complete, found " + strconv.Itoa(len(projects)) + " projects")
	for key, project := range projects {
		project.Name = key // Slightly hacky workaround to get the key from the map assigned to the struct
		log.Println("Project name: " + project.Name + " (" + path.Join(project.ProjectDirectory, project.SqlDirectory) + ")")
	}

	// @todo - dial into each database, update state in migration manager sqlite db
	// @todo - compute list of complete vs pending migrations, add to cache
	// @todo - start http listener
	r := mux.NewRouter()
	r.Use(middleware.JsonContentTypeMiddleware)

	r.HandleFunc("/", handlers.RootHandler).Methods("GET", "OPTIONS")

	log.Println("Listening on port 8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
