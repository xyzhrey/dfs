package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(url string) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), url)
}
