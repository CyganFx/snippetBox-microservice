package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CyganFx/snippetBox-microservice/news/grpc/news_pb"
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type Server struct {
	news_pb.UnimplementedNewsServiceServer
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

//TODO
func (s *Server) CreateNews(ctx context.Context, req *news_pb.NewsCreateRequest) (*news_pb.NewsCreateResponse, error) {
	log.Printf("GetNews function was invoked with %v \n", req)
	//title := req.GetTitle()
	//content := req.GetContent()
	//expires := req.GetExpires()

	panic("implement me")
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()
	news_pb.RegisterNewsServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
