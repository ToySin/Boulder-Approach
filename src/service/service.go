package service

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/toysin/boulder/src/api"
	db "github.com/toysin/boulder/src/client/database"
)

type Service struct {
	pb.UnimplementedBoulderApproachServiceServer

	dbClient *db.Postgres
}

func New(connStr string) (*Service, error) {
	dbClient, err := db.New(connStr)
	if err != nil {
		return nil, err
	}
	return &Service{dbClient: dbClient}, nil
}

func (s *Service) UpsertBoulder(ctx context.Context, req *pb.UpsertBoulderRequest) (*pb.UpsertBoulderResponse, error) {
	boulder := &pb.Boulder{
		Id:   uuid.NewString(),
		Name: req.Name,
	}
	err := s.dbClient.SaveBoulder(boulder)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UpsertBoulderResponse{}, nil
}

func (s *Service) UpsertApproach(ctx context.Context, req *pb.UpsertApproachRequest) (*pb.UpsertApproachResponse, error) {
	// TODO: Get the boulder data from DB
	boulder, err := s.dbClient.GetBoulder(req.Boulder.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// Build Approach from xml data
	gpxData, err := ParseGPX(req.GpxXml)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	points := []*pb.Point{}
	channels := make([]chan []*pb.Point, 1)
	for _, track := range gpxData.Tracks {
		for _, segment := range track.Segments {
			ch := make(chan []*pb.Point)
			channels = append(channels, ch)
			go func() {
				points := []*pb.Point{}
				for _, point := range segment.Points {
					p := &pb.Point{
						Latitude:  point.Latitude,
						Longitude: point.Longitude,
					}
					points = append(points, p)
				}
				ch <- points
			}()
		}
	}

	for _, ch := range channels {
		points = append(points, <-ch...)
	}

	approach := &pb.Approach{
		Id:          uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
		Points:      points,
	}

	// Save the approach to DB
	err = s.dbClient.SaveApproach(req.BoulderId, approach)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpsertApproachResponse{}, nil
}
