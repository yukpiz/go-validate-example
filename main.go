package main

import (
	context "context"
	fmt "fmt"
	"net"

	"github.com/yukpiz/go-validate-example/pb"
	grpc "google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	h := &handler{}
	pb.RegisterExampleServer(server, h)

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

func (*handler) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, newError(ctx, err)
	}
	return &pb.HelloResponse{
		Id:        req.Id,
		LastName:  req.LastName,
		FirstName: req.FirstName,
	}, nil
}

func newError(ctx context.Context, err error) error {
	return err
}
