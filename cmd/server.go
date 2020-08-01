package main

import (
	"context"
	"github.com/marc47marc47/grpc-protobuf/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Add(c context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}
func (s *server) Multiply(c context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}
func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
