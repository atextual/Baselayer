package model

import (
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	Db *sqlx.DB
}
