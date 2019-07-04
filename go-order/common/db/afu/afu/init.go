package afu

import "github.com/jmoiron/sqlx"

type AfuBbModel struct {
	Db *sqlx.DB
}
