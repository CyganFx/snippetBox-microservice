package grpc_client

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"log"
	"snippetBox-microservice/catalog/api/grpc/protobuffs"
	"strconv"
	"time"
)

func DoGetNews(c protobuffs.NewsServiceClient, id int32) *protobuffs.NewsGetResponse {
	ctx := context.Background()
	request := &protobuffs.NewsGetRequest{Id: id}

	response, err := c.SendNews(ctx, request)
	if err != nil {
		log.Fatalf("error while calling GetNews %v", err)
	}
	log.Printf("response from GetNews: id: %v, title: %s, content: %s, created: %v, expires: %v  ",
		response.Id, response.Title, response.Content, response.Created, response.Expires)

	return response
}

func DoCreateNews(c protobuffs.NewsServiceClient, title, content, expires string) *protobuffs.NewsCreateResponse {
	ctx := context.Background()
	integerExpires, err := strconv.Atoi(expires)
	if err != nil {
		log.Fatal(err)
	}
	timeTimeExpires := time.Now().AddDate(0, 0, integerExpires)

	expiresTimeTimestamp, err := ptypes.TimestampProto(timeTimeExpires)
	if err != nil {
		log.Fatal("Failed to convert expires time to timestamp")
	}

	request := &protobuffs.NewsCreateRequest{
		Title:   title,
		Content: content,
		Expires: expiresTimeTimestamp,
	}

	response, err := c.CreateNews(ctx, request)
	if err != nil {
		log.Fatalf("error while calling CreateNews %v", err)
	}

	log.Printf("response from CreateNews: id: %v",
		response.Id)

	return response
}
