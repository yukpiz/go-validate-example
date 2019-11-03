package main

import (
	context "context"
	fmt "fmt"
	"net"

	grpc "google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	h := &handler{}
	RegisterExampleServer(server, h)

	p, err := net.Listen("tcp", ":1111")
	if err != nil {
		panic(err)
	}

	fmt.Println("Start Server")
	if err := server.Serve(p); err != nil {
		panic(err)
	}
}

type handler struct{}

func (*handler) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, nil
}
