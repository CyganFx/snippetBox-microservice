package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"snippetBox-microservice/catalog/api/grpc/protobuffs"
	"snippetBox-microservice/catalog/internal/repository"
	"snippetBox-microservice/catalog/pkg/domain"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Server struct {
	protobuffs.UnimplementedCatalogServiceServer
	catalogRepository repository.ICatalogRepository
}

func (s *Server) SendProduct(ctx context.Context, req *protobuffs.ProductSendRequest) (*protobuffs.ProductSendResponse, error) {
	log.Printf("SendProduct function was invoked with %v \n", req)
	id := req.GetId()

	url := fmt.Sprintf("https://localhost:4012/catalog/%v", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to get product: %v", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	product := &domain.Product{}
	json.NewDecoder(resp.Body).Decode(product)
	json.Unmarshal(body, &product)

	result := &protobuffs.ProductSendResponse{
		Id:          int32(product.ID),
		Title:       product.Title,
		Category:    product.Category,
		Description: product.Description,
		Price:       product.Price,
	}

	return result, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	grpcServer := grpc.NewServer()

	dsn := flag.String("dsn",
		os.Getenv("db_url_grpc"),
		"PostgreSQL data source name")
	flag.Parse()

	dbPool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	catalogRepository := repository.CatalogRepository{Pool: dbPool}

	protobuffs.RegisterCatalogServiceServer(grpcServer, &Server{catalogRepository: catalogRepository})
	log.Println("Server is running on port:50052")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
