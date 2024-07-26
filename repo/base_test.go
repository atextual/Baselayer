package repo

import (
	"BaseLayer/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func MockConnection() (*model.Connection, sqlmock.Sqlmock) {
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	return &model.Connection{
		Db: sqlx.NewDb(mockdb, "sqlmock"),
	}, mock
}
