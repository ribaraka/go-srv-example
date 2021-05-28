package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ribaraka/go-srv-example/pkg/models"
)

type EmailTokenRepository struct {
	pool *pgxpool.Pool
}

func NewVerificationRepository(pxPool *pgxpool.Pool) *EmailTokenRepository {
	return &EmailTokenRepository{
		pool: pxPool,
	}
}

func (mr *EmailTokenRepository) GetByToken(ctx context.Context, token string) (*models.EmailVerificationToken, error) {
	user := &models.EmailVerificationToken{}
	err := mr.pool.QueryRow(ctx,
		`SELECT * FROM email_verification_tokens WHERE verification_token=$1`, token).Scan(&user.UserId, &user.VerificationToken, &user.GeneratedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}