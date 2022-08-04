package grpc

import (
	"context"
	pb "go-news-clean/internal/proto"
	"log"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type contentCheckServiceServer struct {
	pb.UnimplementedContentCheckServiceServer
}

func (*contentCheckServiceServer) CheckHealth(ctx context.Context, r *pb.EmptyRequest) (*pb.HealthResponse, error) {

	log.Printf("[REQ] CheckHealth")

	if service_alive() {
		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "200"))
		return &pb.HealthResponse{
			ServiceName:   "ContentCheckService",
			ServiceStatus: "200",
		}, nil
	} else {
		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "500"))
		return &pb.HealthResponse{
			ServiceName:   "ContentCheckService",
			ServiceStatus: "500",
		}, nil
	}
}

func service_alive() bool {
	return (rand.Intn(100) < 50)
	// return true
}
