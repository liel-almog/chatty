package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreSQLpgx struct {
	pool *pgxpool.Pool
}

var db *PostgreSQLpgx

func Init() {
	db, _ = newPostgreSQLpgx()
}

func newPostgreSQLpgx() (*PostgreSQLpgx, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQLpgx{
		pool: pool,
	}, nil
}

func (p *PostgreSQLpgx) Close() {
	p.pool.Close()
}

func GetDB() *PostgreSQLpgx {
	return db
}
