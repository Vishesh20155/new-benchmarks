package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	pb "github.com/vishesh20155/new-benchmarks/gen/geo"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeoServer
}

func (s *server) Nearby(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	var unserializedReq pb.UnserializedRequest
	err := json.Unmarshal([]byte(req.Req), &unserializedReq)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %v", err)
	}

	// Implement your business logic here
	// For this example, we'll just return some dummy data
	hotelIds := []string{"hotel1", "hotel2", "hotel3"}

	unserializedResp := &pb.UnserializedResponse{
		HotelIds: hotelIds,
	}

	respBytes, err := json.Marshal(unserializedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %v", err)
	}

	return &pb.Result{
		Resp: string(respBytes),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGeoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
