package database

import (
	"database/sql"

	"github.com/arjendevos/instagram-client/helpers"
	_ "github.com/lib/pq"
)

func StartDB() (*sql.DB, error) {
	postgresqlUrl := helpers.GetEnv("POSTGRESQL_URL")

	return sql.Open("postgres", postgresqlUrl)
}
