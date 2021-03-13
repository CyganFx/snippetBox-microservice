package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/CyganFx/snippetBox-microservice/user_details/pkg/repository"
	"github.com/CyganFx/snippetBox-microservice/user_details/pkg/service"
	"github.com/golangcollege/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
	templateCache  map[string]*template.Template
	userService    service.UserServiceInterface
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// not using cookie store from oauth, so ignore errors in console
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			os.Getenv("CALLBACK_URL")),
	)
}

func main() {
	addr := flag.String("addr", ":4002", "HTTP network address")
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

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	userRepository := repository.NewUserRepository(dbPool)

	userService := service.NewUserService(userRepository)

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		session:        session,
		userService:    userService,
		templateCache:  templateCache,
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
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	//err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
