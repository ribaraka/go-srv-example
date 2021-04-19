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

func SQLStatements(user models.User) {
	config, err := conf.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx := context.Background()
	dbpool, err := pgxpool.Connect(ctx, config.DBSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	sqlAddUser := `INSERT INTO users (firstName, lastName, email) VALUES ($1, $2, $3)`
	_, err = dbpool.Exec(ctx, sqlAddUser, user.FirstName, user.LastName, user.Email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database:: %v\n", err)
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

/*func OpenConnection() *sql.DB {

	config, err := conf.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
*/