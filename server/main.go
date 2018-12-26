package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/gidyon/microservices/grpc/hello"
)

type helloServer struct{}

func (hs *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := &pb.HelloResponse{Reply: fmt.Sprintf("Hello %q, we have received your message", req.Name)}
	return reply, nil
}

func main() {
	srv := grpc.NewServer()
	pb.RegisterHelloServiceServer(srv, &helloServer{})

	l, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalf("could not listen on port :7070 : %v\n", err)
	}

	log.Println("gRPC server started on port :7070 ...")
	err = srv.Serve(l)
	if err != nil {
		log.Fatalf("could not serve connections: %v\n", err)
	}
}
