package csdb

import (
  "os"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Open() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    log.Printf("Error connectint to database: %q", err)
  }
  return db
}
