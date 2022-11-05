package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/multischema/ent"
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
	http.HandleFunc("/api/v1/auth", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "auth token")
	})
	http.HandleFunc("/api/v1/refresh", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "refresh token")
	})
}
