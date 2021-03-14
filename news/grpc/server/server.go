package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/CyganFx/snippetBox-microservice/news/cmd/handler"
	"github.com/CyganFx/snippetBox-microservice/news/cmd/helpers"
	"github.com/CyganFx/snippetBox-microservice/news/grpc/news_pb"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/repository"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/service"
	"github.com/golang/protobuf/ptypes"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type Server struct {
	news_pb.UnimplementedNewsServiceServer
	newsHandler handler.NewsHandlerInterface
}

// TODO Test
func (s *Server) GetNews(ctx context.Context, req *news_pb.NewsGetRequest) (*news_pb.NewsGetResponse, error) {
	log.Printf("GetNews function was invoked with %v \n", req)
	id := req.GetId()

	url := fmt.Sprintf("https://localhost:4001/news/%v", id)
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

	result := &news_pb.NewsGetResponse{
		Id:      int32(news.ID),
		Title:   news.Title,
		Content: news.Content,
		Created: createdTime,
		Expires: expiresTime,
	}
	return result, nil
}

//TODO test
func (s *Server) CreateNews(ctx context.Context, req *news_pb.NewsCreateRequest) (*news_pb.NewsCreateResponse, error) {
	log.Printf("GetNews function was invoked with %v \n", req)
	title := req.GetTitle()
	content := req.GetContent()
	expires := req.GetExpires().String()

	id, errors := s.newsHandler.CreateNews(title, content, expires)
	if errors != nil {
		for _, e := range errors {
			return nil, fmt.Errorf(e)
		}
	}

	result := &news_pb.NewsCreateResponse{Id: int32(id)}

	return result, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()

	dsn := flag.String("dsn",
		os.Getenv("db_url"),
		"PostgreSQL data source name")

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	dbPool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbPool.Close()

	newsRepository := repository.NewNewsRepository(dbPool)
	newsService := service.NewNewsService(newsRepository)
	helper := helpers.New(errorLog)
	newsHandler := handler.New(newsService, helper)

	news_pb.RegisterNewsServiceServer(s, &Server{newsHandler: newsHandler})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
