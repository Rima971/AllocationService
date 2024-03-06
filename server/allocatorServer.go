package main

import (
	pb "github.com/rima971/allocator/allocator"
	"github.com/rima971/allocator/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":8090"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}
	serverRegistrar := grpc.NewServer()

	c := service.NewDeliveryAgentAllocatorService()

	pb.RegisterDeliveryAgentAllocatorServiceServer(serverRegistrar, c)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %v", err)
	}
}
