package postgres

import (
	"database/sql"
	"log"

	"github.com/gopheramol/domain-driven-arch/config"

	_ "github.com/lib/pq"
)

// Connect establishes a connection to PostgreSQL and returns a DB instance
func Connect(config config.Config) (*sql.DB, error) {

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL!")

	return db, nil
}
