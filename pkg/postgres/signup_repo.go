package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ribaraka/go-srv-example/config"
	"github.com/ribaraka/go-srv-example/pkg/crypto"
	"github.com/ribaraka/go-srv-example/pkg/email"
	"github.com/ribaraka/go-srv-example/pkg/models"
)

type SignUpRepository struct {
	pool *pgxpool.Pool
}

func NewSignUpRepository(pxPool *pgxpool.Pool) *SignUpRepository {
	return &SignUpRepository{
		pool: pxPool,
	}
}

func (sr *SignUpRepository) SQLStatements(ctx context.Context, user models.User, c config.Config) error {
	insertUser := `INSERT INTO users (firstName, lastName, email) VALUES ($1, $2, $3)`
	_, err := sr.pool.Exec(ctx, insertUser, user.FirstName, user.LastName, user.Email)
	if err != nil {
		return fmt.Errorf("Unable to insert data into database:: %v\n", err)
	}

	hash, err := crypto.HashAndSalt([]byte(user.Password))
	if err != nil {
		return fmt.Errorf("failed to hash crypto: %w", err)
	}
	insertPasswordHash := `INSERT INTO credentials (password_hash) VALUES ($1)`
	_, err = sr.pool.Exec(ctx, insertPasswordHash, hash)
	if err != nil {
		return fmt.Errorf("Unable to insert hash into password_hash:: %v\n", err)
	}

	emailToken := crypto.GenerateToken(32)
	insertEmailToken := `INSERT INTO email_verification_tokens (verification_token) VALUES ($1)`
	_, err = sr.pool.Exec(ctx, insertEmailToken, emailToken)
	if err != nil {
		return fmt.Errorf("Unable to insert token: %v\n", err)
	}

	err = email.SendVerifyMassage(c, user.Email, emailToken)
	if err != nil{
		return fmt.Errorf("Unable to send email %v\n", err)
	}

	return nil
}

func (sr *SignUpRepository) GetByID(ctx context.Context, id int) (*models.TableUser, error) {
	user := &models.TableUser{}
	err := sr.pool.QueryRow(ctx,
		`SELECT * FROM users WHERE id=$1`, id).Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Verified)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (sr *SignUpRepository) UpdateUserByEmail(ctx context.Context, user *models.TableUser) error {
	_, err := sr.pool.Exec(ctx,
		`UPDATE users SET firstname = $2, lastname = $3, email = $4, verified = $5 WHERE email = $1`,
	user.Email, user.Firstname, user.Lastname, user.Email, user.Verified)
	if err != nil {
		return fmt.Errorf("Unable to update row: %v\n", err)
	}

	return nil
}