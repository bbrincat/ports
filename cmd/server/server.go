package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "ports/ports"

)

const (
	address     = "localhost:50051"
	defaultName = "world"
)



const (
	port = ":50051"
)

//TODO implement saving to database.

type server struct {
	pb.UnimplementedPortDomainServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) AddPort(ctx context.Context, in *pb.AddPortRequest) (*pb.AddPortReply, error) {
	log.Printf("Received: %v", in.GetCode())
	return &pb.AddPortReply{} ,nil
}

func (s *server) GetPort(ctx context.Context, in *pb.PortRequest) (*pb.GetPortReply, error) {
	log.Printf("Received: %v", in.GetCode())
	return &pb.GetPortReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPortDomainServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//
}
