package configuration

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
)

func NewDatabase(config Config) *sql.DB {
	username := config.Get("DB_USERNAME")
	password := config.Get("DB_PASSWORD")
	host := config.Get("DB_HOST")
	port := config.Get("DB_PORT")
	dbName := config.Get("DB_NAME")
	var dsn string
	if config.Get("ENV") == "production" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-full&sslrootcert=ap-southeast-1-bundle.pem&timezone=UTC", username, password, host, port, dbName)
	} else {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Info("Connected to database.")
	return db
}
