package postgres

import (
	"context"
	"fmt"
	"main/database"
	"os"

	"github.com/jackc/pgx/v5"
)

type NamedArgs = pgx.NamedArgs
type Batch = pgx.Batch
type Identifier = pgx.Identifier
type Client = any
type Database struct{
	Connection *pgx.Conn
	CopyFromSlice func(length int, next func(int) ([]any, error)) pgx.CopyFromSource
	CopyFromRows func(rows [][]any) pgx.CopyFromSource
}

type PostgresDb[TDatabase Database, TClient Client] struct{}

func NewPostgresDb() *PostgresDb[Database, Client] {
	return &PostgresDb[Database, Client]{}
}

func (pb *PostgresDb[TDatabase, TClient]) New(
	name string,
	address string,
	username string,
	password string,
) *database.DbHandler[Database, any] {
	//Creates a new Postgresql database connection
	connection, err := pgx.Connect(context.Background(), os.Getenv("DB_DNS"))
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &database.DbHandler[Database, Client]{
		Database: Database{ 
			Connection: connection,
			CopyFromSlice: pgx.CopyFromSlice,
			CopyFromRows: pgx.CopyFromRows,
		},
		Client: nil,
	}
}