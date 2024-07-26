package repo

import (
	"BaseLayer/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

func AddDatabase(connection *sqlx.DB, model *models.Project) (*models.Project, error) {
	stmt := `INSERT INTO databases VALUES (NULL, ?, ?, ?, ?, ?, ?)`
	_, err := connection.Exec(
		stmt,
		model.Name,
		strings.ToUpper(model.Name),
		model.Database.Driver,
		model.Database.Username,
		model.Database.Password,
		model.Database.Port,
	)

	return model, err
}

func ListDatabases(connection *sqlx.DB) {

}
