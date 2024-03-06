package main

import (
	"context"
	pb "github.com/rima971/allocator/allocator"
	"github.com/rima971/allocator/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

func server(ctx context.Context) (pb.DeliveryAgentAllocatorServiceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()

	c := service.NewDeliveryAgentAllocatorService()
	pb.RegisterDeliveryAgentAllocatorServiceServer(baseServer, c)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %s", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %s", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewDeliveryAgentAllocatorServiceClient(conn)

	return client, closer
}

func TestCurrencyConvertor_convert(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *pb.AllocatorResponse
		err error
	}
	tests := map[string]struct {
		in       *pb.AllocatorRequest
		expected expectation
	}{}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.Allocate(ctx, tt.in)
			if err != nil {
				if tt.expected.err == nil || "rpc error: code = Unknown desc = "+tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.GetDeliveryAgent().GetId() != out.GetDeliveryAgent().GetId() ||
					tt.expected.out.GetDeliveryAgent().GetUserId() != out.GetDeliveryAgent().GetUserId() ||
					tt.expected.out.GetDeliveryAgent().GetCurrentLocationPincode() != out.GetDeliveryAgent().GetCurrentLocationPincode() {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}
}
