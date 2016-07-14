package csdb

import (
  "os"
	"database/sql"
	"github.com/lib/pq"
	"log"
)

func Open() *sql.DB {
	url := os.Getenv("DATABASE_URL")
  connection, _ := pq.ParseURL(url)
  connection += " sslmode=require"
	db, err := sql.Open("postgres", connection)
  if err != nil {
    log.Printf("Error connecting to database: %q", err)
  } else {
		log.Printf("Connected to database")
	}
  return db
}
