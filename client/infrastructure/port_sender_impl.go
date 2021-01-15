package infrastructure

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"ports/client/domain"
	pb "ports/ports"
	"time"
)

//Infrastructure package contains logic related to IO/External side effects.

//Stuct which will implement interface defined in domain.
type GRPCPortSenderService struct{}

//TODO move to configuration file
const (
	address     = "localhost:50051"
	defaultName = "portService"
)

// Initailise and dependencies and return service
func  NewGRPCPortSenderService() GRPCPortSenderService{
	//Potentially initilise here
	return GRPCPortSenderService{}
}

//Implement  interface
func (*GRPCPortSenderService) SendPort(domain.Port){
	// Set up a connection to the server.

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
	log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPortDomainServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
	name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddPort(ctx, &pb.AddPortRequest{Name: &name})
	if err != nil {
	log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatusCode())

}

