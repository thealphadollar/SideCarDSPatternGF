package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/EwanValentine/distributed-patterns/sidecar-http-grpc/app/transport"
	grpc "google.golang.org/grpc"
)

type server struct{}

func (s *server) FetchGreeting(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Reply: "Hello, " + req.Name}, nil
}

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterApplicationServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
