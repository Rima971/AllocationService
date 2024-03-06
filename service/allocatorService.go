package service

import (
	"context"
	"errors"
	pb "github.com/rima971/allocator/allocator"
)

type DeliveryAgentAllocatorService struct {
	pb.UnimplementedDeliveryAgentAllocatorServiceServer
}

func NewDeliveryAgentAllocatorService() *DeliveryAgentAllocatorService {
	return &DeliveryAgentAllocatorService{}
}

func fetchAllDeliveryAgents(pincode int32) ([]pb.DeliveryAgent, error) {
	return nil, errors.New("fetch request for delivery agents failed")
}

func (d *DeliveryAgentAllocatorService) Allocate(ctx context.Context, in *pb.AllocatorRequest) (*pb.AllocatorResponse, error) {
	res, err := fetchAllDeliveryAgents(in.GetDeliveryLocationPincode())
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("No delivery agent available at the moment at the given location")
	}

	da := res[0]

	return &pb.AllocatorResponse{DeliveryAgent: &da}, nil
}
