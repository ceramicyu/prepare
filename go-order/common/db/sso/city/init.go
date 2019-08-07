package city

import (
	"github.com/jmoiron/sqlx"
)

type SsoCityModel struct {
	Db *sqlx.DB
	City *CityDb
}

func NewSsoCityModel(db *sqlx.DB)SsoCityModel{
	c := SsoCityModel{
		Db: db,
	}
	c.City=&CityDb{c}

	return c
}
