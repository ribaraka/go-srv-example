package postgres

import (
	"context"
	"fmt"
	"github.com/ribaraka/go-srv-example/internal/conf"
	"github.com/ribaraka/go-srv-example/pkg/models"
	"github.com/ribaraka/go-srv-example/pkg/password"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// TODO: context should be passed from request
func SQLStatements(user models.User) {

	// TODO: pgxpool.Connect should be executed only once in main function
	dbpool := OpenConnection()
	ctx := context.Background()

	sqlAddUser := `INSERT INTO users (firstName, lastName, email) VALUES ($1, $2, $3)`
	_, err := dbpool.Exec(ctx, sqlAddUser, user.FirstName, user.LastName, user.Email)
	if err != nil {
		// TODO: return fmt.Errorf( "Unable to insert data into database: %w", err)
		fmt.Fprintf(os.Stderr, "Unable to insert data into database:: %v\n", err)
		// TODO: we should not invoke os.Exit
		os.Exit(1)
	}

	pwd := []byte(user.Password)
	hash := password.HashAndSalt(pwd)

	sqlAddCredential := `INSERT INTO credentials (password_hash) VALUES ($1)`
	_, err = dbpool.Exec(ctx, sqlAddCredential, hash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert hash into password_hash:: %v\n", err)
		os.Exit(1)
	}
}

func OpenConnection() *pgxpool.Pool {
	config, err := conf.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	pool, err := pgxpool.Connect(context.Background(), config.DBURL)
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer pool.Close()
	return pool
}
