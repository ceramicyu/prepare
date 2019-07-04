package oa

import "github.com/jmoiron/sqlx"

type SsoOaModel struct {
	Db *sqlx.DB
}
