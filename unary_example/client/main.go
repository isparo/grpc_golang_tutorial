package main

import (
	"context"
	"fmt"

	"github.com/josue/grpc_golang_tutorial/unary_example/server/protofiles/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	res, err := c.Greet(context.Background(), &greetpb.GreetRequest{
		Greeting: "hello",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
