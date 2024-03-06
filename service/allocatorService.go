package service

import (
	"context"
	"errors"
	"github.com/rima971/allocator/adapter"
	pb "github.com/rima971/allocator/allocator"
)

type DeliveryAgentAllocatorService struct {
	pb.UnimplementedDeliveryAgentAllocatorServiceServer
}

func NewDeliveryAgentAllocatorService() *DeliveryAgentAllocatorService {
	return &DeliveryAgentAllocatorService{}
}

func (d *DeliveryAgentAllocatorService) Allocate(ctx context.Context, in *pb.AllocatorRequest) (*pb.AllocatorResponse, error) {
	res, err := adapter.FetchNearestDeliveryAgents(in.GetDeliveryLocationPincode())
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("no delivery agent available at the moment at the given location")
	}

	da := res[0]

	return &pb.AllocatorResponse{DeliveryAgent: &da}, nil
}
