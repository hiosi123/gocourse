package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hiosi123/gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error {
	log.Println("ListBlog was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Error(
			codes.Internal,
			fmt.Sprintf("unkown internal error: %v", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Error(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MonggoDB: %v", err),
			)

		}

		stream.Send(documentToBlog(data))
	}
	if err = cur.Err(); err != nil {
		return status.Error(
			codes.Internal,
			fmt.Sprintf("unkown internal error: %v", err),
		)
	}

	return nil
}
