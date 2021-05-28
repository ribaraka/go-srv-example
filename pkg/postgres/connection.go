package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ribaraka/go-srv-example/config"
)

func OpenConnection(c config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), c.DBURL)
	if err != nil {
		return nil, fmt.Errorf("Unable to connection to database: %v\n", err)
	}

	return pool, nil
}

