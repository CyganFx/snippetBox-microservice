package grpc_server

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"log"
	"snippetBox-microservice/news/api/controller"
	"snippetBox-microservice/news/api/grpc/protobuffs"
	"snippetBox-microservice/news/pkg/domain"
	"time"
)

type Server struct {
	protobuffs.UnimplementedNewsServiceServer
	NewsService controller.NewsServiceInterface
}

func (s *Server) SendNews(ctx context.Context, req *protobuffs.NewsGetRequest) (*protobuffs.NewsGetResponse, error) {
	log.Printf("SendNews function was invoked with %v \n", req)
	id := int(req.GetId())

	news, err := s.NewsService.FindById(id)
	if err != nil {
		return nil, err
	}

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

	id, err := s.NewsService.Save(news)
	if err != nil {
		return nil, err
	}

	result := &protobuffs.NewsCreateResponse{Id: int32(id)}

	return result, nil
}
