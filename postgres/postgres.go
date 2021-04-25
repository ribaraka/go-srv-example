package postgres

import (
	"context"
	"fmt"
	"github.com/ribaraka/go-srv-example/internal/conf"
	"github.com/ribaraka/go-srv-example/pkg/models"
	"github.com/ribaraka/go-srv-example/pkg/password"

	"github.com/jackc/pgx/v4/pgxpool"
)

type SignUpRepository struct {
	pool *pgxpool.Pool
}

func NewSignUpRepository(pool *pgxpool.Pool) *SignUpRepository {
	return &SignUpRepository{
		pool: pool,
	}
}

func (sr *SignUpRepository)SQLStatements(ctx context.Context, user models.User) error {
	sqlAddUser := `INSERT INTO users (firstName, lastName, email) VALUES ($1, $2, $3)`
	_, err := sr.pool.Exec(ctx, sqlAddUser, user.FirstName, user.LastName, user.Email)
	if err != nil {
		return 	fmt.Errorf( "Unable to insert data into database:: %v\n", err)
	}

	hash, err := password.HashAndSalt([]byte(user.Password))

	sqlAddCredential := `INSERT INTO credentials (password_hash) VALUES ($1)`
	_, err = sr.pool.Exec(ctx, sqlAddCredential, hash)
	if err != nil {
		return fmt.Errorf("Unable to insert hash into password_hash:: %v\n", err)

	}
	return nil
}

func OpenConnection(c conf.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), c.DBURL)
	if err != nil {
		return nil, fmt.Errorf("Unable to connection to database: %v\n", err)
	}

	return pool, nil
}
