package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"snippetBox-microservice/catalog/api/grpc/protobuffs"
	"snippetBox-microservice/catalog/internal/controller"
	"snippetBox-microservice/catalog/internal/repository"
	"snippetBox-microservice/catalog/utils/helpers"
	"time"
)

type application struct {
	errorLog          *log.Logger
	infoLog           *log.Logger
	catalogController controller.ICatalogController
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	addr := flag.String("addr", ":4012", "HTTP network address")
	dsn := flag.String("dsn",
		os.Getenv("db_url"),
		"PostgreSQL data source name")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	dbPool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbPool.Close()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	grpcClient := protobuffs.NewNewsServiceClient(conn)

	newsRepository := repository.NewCatalogRepository(dbPool)
	helper := helpers.New(errorLog)
	newsHandler := controller.New(newsRepository, helper, grpcClient)

	app := &application{
		errorLog:          errorLog,
		infoLog:           infoLog,
		catalogController: newsHandler,
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting  server on %v", *addr)
	err = srv.ListenAndServeTLS("./crypto/tls/cert.pem", "./crypto/tls/key.pem")
	errorLog.Fatal(err)
}
