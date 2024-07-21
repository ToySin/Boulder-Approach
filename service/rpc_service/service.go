package rpc_service

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/toysin/boulder/service/api"
)

type Service struct {
	pb.UnimplementedBoulderApproachServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetApproach(
	ctx context.Context,
	req *pb.GetApproachRequest,
) (*pb.GetApproachResponse, error) {
	return &pb.GetApproachResponse{
		Approach: &pb.Approach{
			ApproachId:  "밤바위1",
			Name:        "밤바위",
			Description: "밤바위 어프로치 1",
			Points: []*pb.Point{
				{
					Latitude:  37.35370983373127,
					Longitude: 126.93727609713277,
					Text:      "시작점",
				},
				{
					Latitude:  37.35370669051758,
					Longitude: 126.9373087027361,
					Text:      "1번째 포인트",
				},
				{
					Latitude:  37.353701326099554,
					Longitude: 126.93733887758752,
					Text:      "2번째 포인트",
				},
				{
					Latitude:  37.3536795338507,
					Longitude: 126.93737229421913,
					Text:      "3번째 포인트",
				},
				{
					Latitude:  37.35369578165948,
					Longitude: 126.93741432508976,
					Text:      "4번째 포인트",
				},
				{
					Latitude:  37.3537103859863,
					Longitude: 126.93743402871773,
					Text:      "5번째 포인트",
				},
				{
					Latitude:  37.35371272593596,
					Longitude: 126.93745788058017,
					Text:      "6번째 포인트",
				},
				{
					Latitude:  37.353729733047416,
					Longitude: 126.93746648423485,
					Text:      "종료점",
				},
			},
		},
	}, nil
}

func (s *Service) CreateApproach(
	ctx context.Context,
	req *pb.CreateApproachRequest,
) (*pb.CreateApproachResponse, error) {

	// TODO: Get the boulder data from DB
	// boulder := dbClient.GetBoulder(req.BoulderId)

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
		ApproachId:  uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Points:      points,
	}

	// Save the approach to DB
	// dbClient.SaveApproach(req.BoulderId, points)

	return &pb.CreateApproachResponse{}, nil
}
