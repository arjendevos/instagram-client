package database

import (
	"context"
	"database/sql"

	"github.com/arjendevos/instagram-client/helpers"
	_ "github.com/lib/pq"
)

func startDB() (*sql.DB, error) {
	postgresqlUrl := helpers.GetEnv("POSTGRESQL_URL")

	return sql.Open("postgres", postgresqlUrl)
}

type DatabaseClient struct {
	Ctx context.Context
	DB  *sql.DB
}

func NewDatabaseClient(ctx context.Context) (*DatabaseClient, error) {
	db, err := startDB()
	if err != nil {
		return nil, err
	}
	return &DatabaseClient{Ctx: ctx, DB: db}, nil
}
