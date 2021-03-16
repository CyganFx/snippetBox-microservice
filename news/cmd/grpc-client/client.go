package main

import (
	"context"
	"fmt"
	"github.com/CyganFx/snippetBox-microservice/news/api/grpc/protobuffs"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := protobuffs.NewNewsServiceClient(conn)

	doGetNews(c, 1)
	//doCreateNews(c)
}

func doGetNews(c protobuffs.NewsServiceClient, id int32) {
	ctx := context.Background()
	request := &protobuffs.NewsGetRequest{Id: id}

	response, err := c.GetNews(ctx, request)
	if err != nil {
		log.Fatalf("error while calling GetNews %v", err)
	}
	log.Printf("response from GetNews: id: %v, title: %s, content: %s, created: %v, expires: %v  ",
		response.Id, response.Title, response.Content, response.Created, response.Expires)
}

func doCreateNews(c protobuffs.NewsServiceClient) {
	title := "some title"
	content := "some content"
	expires := "7"

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
}
