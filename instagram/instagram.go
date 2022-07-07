package instagram

import (
	"context"
	"database/sql"
)

type Instagram struct {
	ctx context.Context
	db  *sql.DB
}

func NewInstagram(ctx context.Context, db *sql.DB) *Instagram {
	return &Instagram{ctx: ctx, db: db}
}
