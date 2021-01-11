package main

import (
	"alexedwards.net/snippetbox/pkg/models/psql"
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *psql.SnippetModel
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	os.Setenv("dsn", "user=web password=duman070601 "+
		"host=localhost port=5432 dbname=snippetbox pool_max_conns=10")

	dsn := flag.String("dsn",
		os.Getenv("dsn"),
		"PostgreSQL data source name")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	dbPool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &psql.SnippetModel{DB: dbPool},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting  server on %v", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
