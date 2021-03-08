package main

// HTTPS doesn't work with postman

import (
	"alexedwards.net/snippetbox/pkg/repository"
	"alexedwards.net/snippetbox/pkg/service"
	"context"
	"crypto/tls"
	"flag"
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
	//snippetRepository *repository.SnippetRepositoryInterface
	//userRepository    *repository.UserRepositoryInterface
	templateCache  map[string]*template.Template
	snippetService service.SnippetServiceInterface
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

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	//snippetRepository := repository.NewSnippetRepository(dbPool)
	//userRepository := repository.NewUserRepository(dbPool)
	//
	//snippetService := service.NewSnippetService(snippetRepository)
	//userService := service.NewUserService(userRepository)

	snippetRepository := &repository.SnippetRepository{Pool: dbPool}
	userRepository := &repository.UserRepository{Pool: dbPool}

	snippetService := &service.SnippetService{SnippetRepository: snippetRepository}
	userService := &service.UserService{UserRepository: userRepository}

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		session:        session,
		snippetService: snippetService,
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
	errorLog.Fatal(err)
}
