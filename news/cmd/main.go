package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/CyganFx/snippetBox-microservice/news/internal/controller"
	"github.com/CyganFx/snippetBox-microservice/news/internal/repository"
	"github.com/CyganFx/snippetBox-microservice/news/internal/service"
	"github.com/CyganFx/snippetBox-microservice/news/utils/helpers"
	"github.com/golangcollege/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	session     *sessions.Session
	newsHandler controller.NewsControllerInterface
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	addr := flag.String("addr", ":4001", "HTTP network address")
	dsn := flag.String("dsn",
		os.Getenv("db_url"),
		"PostgreSQL data source name")
	secret := flag.String("secret", os.Getenv("SESSION_SECRET"), "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	dbPool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbPool.Close()

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	newsRepository := repository.NewNewsRepository(dbPool)
	newsService := service.NewNewsService(newsRepository)
	helper := helpers.New(errorLog)
	newsHandler := controller.New(newsService, helper)

	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		session:     session,
		newsHandler: newsHandler,
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
