package main

// HTTPS doesn't work with postman

//import (
//	"alexedwards.net/snippetbox/pkg/repository"
//	"context"
//	"crypto/tls"
//	"flag"
//	"github.com/golangcollege/sessions"
//	"github.com/jackc/pgx/v4/pgxpool"
//	"github.com/joho/godotenv"
//	"html/template"
//	"log"
//	"net/http"
//	"os"
//	"time"
//)
//
//type application struct {
//	errorLog      *log.Logger
//	infoLog       *log.Logger
//	session       *sessions.Session
//	snippets      *repository.SnippetModel
//	users         *repository.UserModel
//	templateCache map[string]*template.Template
//}
//
//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}
//
//func main() {
//	addr := flag.String("addr", ":4001", "HTTP network address")
//	dsn := flag.String("dsn",
//		os.Getenv("db_url"),
//		"PostgreSQL data source name")
//	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
//	flag.Parse()
//
//	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
//	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
//
//	dbPool, err := pgxpool.Connect(context.Background(), *dsn)
//	if err != nil {
//		errorLog.Fatal(err)
//	}
//	defer dbPool.Close()
//
//	templateCache, err := newTemplateCache("./ui/html/")
//	if err != nil {
//		errorLog.Fatal(err)
//	}
//
//	session := sessions.New([]byte(*secret))
//	session.Lifetime = 12 * time.Hour
//	session.Secure = true
//
//	app := &application{
//		errorLog:      errorLog,
//		infoLog:       infoLog,
//		session:       session,
//		snippets:      &repository.SnippetModel{DB: dbPool},
//		users:         &repository.UserModel{DB: dbPool},
//		templateCache: templateCache,
//	}
//
//	// Initialize a tls.Config struct to hold the non-default TLS settings we want
//	// the server to use.
//	tlsConfig := &tls.Config{
//		PreferServerCipherSuites: true,
//		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
//	}
//
//	srv := &http.Server{
//		Addr:     *addr,
//		ErrorLog: errorLog,
//		Handler:  app.routes(),
//		TLSConfig: tlsConfig,
//		//Add Idle, Read and Write timeouts to the server.
//		IdleTimeout:  time.Minute, //do not delete!
//		ReadTimeout:  5 * time.Second,
//		WriteTimeout: 10 * time.Second,
//	}
//
//	infoLog.Printf("Starting  server on %v", *addr)
//	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
//	errorLog.Fatal(err)
//}
