package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"auth/ent"
	"auth/server"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new connection
func Open(uri string) *ent.Client {
	db, err := sql.Open("pgx", uri)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

var client *ent.Client

func main() {
	client = Open("postgres://postgres:password@auth-db/postgres?sslmode=disable")
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("error: %v", err)
	}
	secret, _ := server.ReadSecret("../certs/private.key")
	s := server.NewAuthServer(context.Background(), client, ":8080", secret)
	log.Fatal(s.ListenAndServeTLS("../certs/localhost-key.pem", "../certs/localhost.pem"))
}
