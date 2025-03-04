package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-gateway-funnzy-dev/pbs"

	"google.golang.org/grpc"
)

type server struct {
}

func (server) Hello(ctx context.Context, helloRequest *pbs.HelloRequest) (*pbs.HelloResponse, error) {
	message := helloRequest.GetMessage()
	log.Printf("receive message %s\n", message)
	helloResponse := &pbs.HelloResponse{
		Message: message,
	}
	return helloResponse, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	pbs.RegisterHelloServer(s, &server{})

	fmt.Println("demo gateway service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
