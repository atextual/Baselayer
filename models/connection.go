package models

import (
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	Db *sqlx.DB
}
