package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func OpenDB(port, user, password, dbname string) (*sql.DB, error) {
	// postgres://postgres:postgres@localhost:5432/greenlight?sslmode=disable
	dsn := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", user, password, port, dbname)
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	// struct.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Using PingContext() for checking that connection is alive and ok
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected successfuly!!!")

	// Return the sql.DB connection pool.
	return db, nil
}
