package main

import (
	"context"
	"fmt"
	"net"

	"github.com/josue/grpc_golang_tutorial/unary_example/server/protofiles/greetpb"
	"google.golang.org/grpc"
)

// this struct could be used to dependency injection
type server struct{}

// handler
func (s server) Greet(ctx context.Context, gr *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("getting : ", gr.Greeting)

	return &greetpb.GreetResponse{
		Result: gr.Greeting + " mrs",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("starting server")

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, server{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
