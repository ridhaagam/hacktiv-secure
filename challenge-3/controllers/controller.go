package controllers

import "secure/challenge-3/repository/postgres/sqlc"

type PgDB struct {
	Master *sqlc.Queries
}
