package repo

import (
	"BaseLayer/model"
	"errors"
	"github.com/jmoiron/sqlx"
)

func AddDatabase(connection *sqlx.DB, model *model.Project) (*model.Project, error) {
	stmt := `INSERT INTO databases VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := connection.Exec(
		stmt,
		model.Database.Name,
		model.Database.NormalisedName,
		model.Database.Driver,
		model.Database.NormalisedDriver,
		model.Database.Username,
		model.Database.Password,
		model.Database.Database,
		model.Database.Port,
	)

	newId, err := res.LastInsertId()
	model.Database.Id = int(newId)

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

func DeleteDatabase(connection *sqlx.DB, database *model.Database) error {
	if connection == nil {
		cxn, err := GetConnection()
		if err != nil {
			return err
		}

		connection = cxn.Db
	}

	stmt := `DELETE FROM databases WHERE id = ?`
	res, err := connection.Exec(stmt, database.Id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no rows were affected by the delete operation")
	}

	return nil
}
