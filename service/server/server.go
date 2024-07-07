package main

import (
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	pb "github.com/toysin/boulder/service/api"
)

func main() {
	// Create a new server
	s := grpc.NewServer()

	// Register the server with the proto
	pb.RegisterBoulderApproachServiceServer(s, pb.UnimplementedBoulderApproachServiceServer{})

	// Create a listener
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		slog.Error("failed to listen", "reason", err)
	}

	// Serve the listener
	go func() {
		if err := s.Serve(lis); err != nil {
			slog.Error("failed to serve", "reason", err)
		}
	}()
	slog.Info("server started", "address", "localhost:8080")

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	<-sigChannel
	slog.Info("shutting down server")

	s.GracefulStop()
}
