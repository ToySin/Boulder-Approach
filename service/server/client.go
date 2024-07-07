package main

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/toysin/boulder/service/api"
)

func main() {
	// Create a new client
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("failed to connect", "reason", err)
	}
	client := pb.NewBoulderApproachServiceClient(conn)
	if err != nil {
		slog.Error("failed to create client", "reason", err)
	}

	// Create a new request
	req := &pb.GetApproachRequest{
		ApproachId: "밤바위1",
	}

	// Call the service
	resp, err := client.GetApproach(context.Background(), req)
	if err != nil {
		slog.Error("failed to call service", "reason", err)
	}

	// Print the response
	slog.Info("response", "response", resp)
}
