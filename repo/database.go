package repo

import (
	"BaseLayer/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

func AddDatabase(connection *sqlx.DB, model *model.Project) (*model.Project, error) {
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

func ListDatabases(connection *sqlx.DB) ([]model.Database, error) {
	var databases []model.Database

	if connection == nil {
		cxn, err := GetConnection()
		if err != nil {
			return nil, err
		}

		connection = cxn.Db
	}

	stmt := `SELECT * FROM databases`
	rows, err := connection.Queryx(stmt)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var database model.Database
		err = rows.StructScan(&database)

		if err != nil {
			return nil, err
		}

		databases = append(databases, database)
	}

	return databases, nil
}

func GetDatabase(connection *sqlx.DB, id int) (*model.Database, error) {
	var database model.Database

	if connection == nil {
		cxn, err := GetConnection()
		if err != nil {
			return nil, err
		}

		connection = cxn.Db
	}

	stmt := `SELECT * FROM databases WHERE id = ?`
	row := connection.QueryRowx(stmt, id)
	err := row.StructScan(&database)
	if err != nil {
		return nil, err
	}

	return &database, nil
}
