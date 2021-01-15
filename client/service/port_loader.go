package service

import (
	"encoding/json"
	"io"
	"log"
	"ports/client/domain"
)

// Implementation of business logic
type GRPCPortLoaderService struct{}

func  NewGRPCPortLoaderService() GRPCPortLoaderService{
	//Potentially initilise here
	return GRPCPortLoaderService{}
}

func (*GRPCPortLoaderService) LoadPorts(stream io.Reader,sender domain.PortSenderService){

	//TODO: Extract Json using decoder and input stream to avoid loading entire file in memory.

	dec := json.NewDecoder(stream)
	for {
		var m map[string]domain.Port
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	//TODO: send json via GRPC Client
}