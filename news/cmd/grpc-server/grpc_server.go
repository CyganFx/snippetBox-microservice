package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"snippetBox-microservice/news/api/grpc/protobuffs"
	"snippetBox-microservice/news/internal/repository"
	"snippetBox-microservice/news/internal/service"
	"snippetBox-microservice/news/pkg/domain"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Server struct {
	protobuffs.UnimplementedNewsServiceServer
	newsService service.NewsServiceInterface
}

//Actually gives news data
func (s *Server) SendNews(ctx context.Context, req *protobuffs.NewsGetRequest) (*protobuffs.NewsGetResponse, error) {
	log.Printf("GetNews function was invoked with %v \n", req)
	id := req.GetId()

	url := fmt.Sprintf("https://localhost:4011/news/%v", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to get news: %v", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	news := &domain.News{}
	json.NewDecoder(resp.Body).Decode(news)
	json.Unmarshal(body, &news)

	createdTime, err := ptypes.TimestampProto(news.Created)
	if err != nil {
		log.Printf("Failed to convert create time to timestamp")
		return nil, err
	}
	expiresTime, err := ptypes.TimestampProto(news.Expires)
	if err != nil {
		log.Printf("Failed to convert expires time to timestamp")
		return nil, err
	}

	result := &protobuffs.NewsGetResponse{
		Id:      int32(news.ID),
		Title:   news.Title,
		Content: news.Content,
		Created: createdTime,
		Expires: expiresTime,
	}
	return result, nil
}

//TODO test
func (s *Server) CreateNews(ctx context.Context, req *protobuffs.NewsCreateRequest) (*protobuffs.NewsCreateResponse, error) {
	log.Printf("CreateNews function was invoked with %v \n", req)
	title := req.GetTitle()
	content := req.GetContent()
	tempExpires := req.GetExpires()
	expires, err := ptypes.Timestamp(tempExpires)
	if err != nil {
		log.Printf("Failed to convert expires time to timestamp")
		return nil, err
	}

	news := &domain.News{
		ID:      -1,
		Title:   title,
		Content: content,
		Created: time.Time{},
		Expires: expires,
	}

	id, err := s.newsService.Save(news)
	if err != nil {
		return nil, err
	}

	result := &protobuffs.NewsCreateResponse{Id: int32(id)}

	return result, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
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

	newsRepository := repository.NewNewsRepository(dbPool)
	newsService := service.NewNewsService(newsRepository)

	protobuffs.RegisterNewsServiceServer(grpcServer, &Server{newsService: newsService})
	log.Println("Server is running on port:50051")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
