package database

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	pb "github.com/toysin/boulder/src/api"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "100days"
	dbname   = "service"
)

var (
	connStr       = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dummyApproach = &pb.Approach{
		Id:          uuid.NewString(),
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
		},
	}
)

func TestConnection(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			_, err := New(connStr)
			if err != nil {
				t.Errorf("New() = %v; want nil", err)
			}
		})
	})
}

func TestUpsertBoulder(t *testing.T) {
	t.Run("UpsertBoulder", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			db, err := New(connStr)
			if err != nil {
				t.Errorf("New() = %v; want nil", err)
			}

			boulder := &pb.Boulder{
				Id:   uuid.NewString(),
				Name: "Test Boulder",
			}

			err = db.SaveBoulder(boulder)
			if err != nil {
				t.Errorf("UpsertBoulder() = %v; want nil", err)
			}

			boulder.Name = "Updated Boulder"

			err = db.SaveBoulder(boulder)
			if err != nil {
				t.Errorf("UpsertBoulder() = %v; want nil", err)
			}
		})
	})
}

func TestUpsertApproach(t *testing.T) {
	t.Run("UpsertApproach", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			db, err := New(connStr)
			if err != nil {
				t.Errorf("New() = %v; want nil", err)
			}

			boulder := &pb.Boulder{
				Id:   uuid.NewString(),
				Name: "Test Boulder",
			}

			err = db.SaveBoulder(boulder)
			if err != nil {
				t.Errorf("UpsertBoulder() = %v; want nil", err)
			}

			err = db.SaveApproach(boulder.Id, dummyApproach)
			if err != nil {
				t.Errorf("UpsertApproach() = %v; want nil", err)
			}

			err = db.SaveApproach(uuid.NewString(), dummyApproach)
			if err == nil {
				t.Errorf("UpsertApproach() = nil; want error")
			}
		})
	})
}

func TestGetBoulder(t *testing.T) {
	t.Run("GetApproach", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			db, err := New(connStr)
			if err != nil {
				t.Errorf("New() = %v; want nil", err)
			}

			boulder := &pb.Boulder{
				Id:   uuid.NewString(),
				Name: "Test Boulder",
			}

			err = db.SaveBoulder(boulder)
			if err != nil {
				t.Errorf("UpsertBoulder() = %v; want nil", err)
			}

			err = db.SaveApproach(boulder.Id, dummyApproach)
			if err != nil {
				t.Errorf("UpsertApproach() = %v; want nil", err)
			}

			findBoulder, err := db.GetBoulder(boulder.Id)
			if err != nil {
				t.Errorf("GetBoulder() = %v; want nil", err)
			}

			if findBoulder.Id != boulder.Id {
				t.Errorf("GetBoulder() = %v; want %v", findBoulder.Id, boulder.Id)
			}
			if len(findBoulder.Approaches) != 1 {
				t.Errorf("GetBoulder() = %v; want 1", len(findBoulder.Approaches))
			}
			if !proto.Equal(findBoulder.Approaches[0], dummyApproach) {
				t.Errorf("GetBoulder() = %v; want %v", findBoulder.Approaches[0], dummyApproach)
			}
		})
	})
}
