package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"snippetBox-microservice/news/api/controller"
	"snippetBox-microservice/news/api/grpc/grpc-server"
	"snippetBox-microservice/news/api/grpc/protobuffs"
	"snippetBox-microservice/news/internal/repository"
	"snippetBox-microservice/news/internal/service"
	"snippetBox-microservice/news/pkg"
	"snippetBox-microservice/news/pkg/rest-errors"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	addr := flag.String("addr", ":4011", "HTTP network address")
	dsn := flag.String("dsn",
		os.Getenv("db_url"),
		"PostgreSQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	dbPool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbPool.Close()

	newsRepository := repository.News(dbPool)
	newsService := service.News(newsRepository)
	restErrorsResponser := rest_errors.NewJsonResponser(errorLog)
	newsController := controller.New(newsService, restErrorsResponser)

	grpcNetListener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	grpcServer := grpc.NewServer()

	protobuffs.RegisterNewsServiceServer(grpcServer, &grpc_server.Server{NewsService: newsService})
	log.Println("GrpcServer is running on port:50051")
	if err := grpcServer.Serve(grpcNetListener); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      pkg.SetupRoutes(newsController),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting  grpc-server on %v", *addr)
	err = srv.ListenAndServeTLS("./crypto/tls/cert.pem", "./crypto/tls/key.pem")
	errorLog.Fatal(err)
}
