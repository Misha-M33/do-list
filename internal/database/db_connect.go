package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	db *pgx.Conn
}

func NewDatabase() *DB {
	return &DB{}
}

func (pg *DB) ConnectDB(user, pass, host, port, name string) error {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)
	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return err
	}
	pg.db = db
	return nil
}
