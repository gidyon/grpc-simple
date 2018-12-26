package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/gidyon/microservices/grpc/hello"
	"google.golang.org/grpc"
)

var (
	name    = flag.String("name", "uknown", "Name of gRPC client")
	message = flag.String("msg", "Hello World", "Message to send to the server")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("127.0.0.1:7070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not dial remote connection: %v", err)
	}
	client := pb.NewHelloServiceClient(conn)

	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: *name, Message: *message})
	if err != nil {
		log.Fatalf("could not get proper response: %v", err)
	}
	log.Println("Server says:", reply.Reply)
}
