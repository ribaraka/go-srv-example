package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ribaraka/go-srv-example/pkg/models"
)

type LoginRepository struct {
	pool *pgxpool.Pool
}

func NewLoginRepository(pxPool *pgxpool.Pool) *LoginRepository {
	return &LoginRepository{
		pool: pxPool,
	}
}

func (mr *LoginRepository) GetByID(ctx context.Context, id int) (*models.Credentials, error) {
	user := &models.Credentials{}
	err := mr.pool.QueryRow(ctx,
		`SELECT * FROM credentials WHERE user_id=$1`, id).Scan(&user.UserId, &user.PasswordHash, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}