package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ribaraka/go-srv-example/internal/conf"
	"log"
)

func OpenConnection() *sql.DB {

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
