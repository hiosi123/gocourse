package main

import (
	"context"
	"log"

	pb "github.com/hiosi123/gRPC/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog has been initiated")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Hongseok2",
		Title:    "this is newly designed",
		Content:  "new design has side effects",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happend while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
