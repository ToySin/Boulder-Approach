package database

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/proto"

	pb "github.com/toysin/boulder/src/api"
)

type Postgres struct {
	db *sql.DB
}

func New(connStr string) (*Postgres, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Postgres{db: db}, nil
}

func (d *Postgres) SaveBoulder(boulder *pb.Boulder) error {
	_, err := d.db.Exec(
		`INSERT INTO boulder (boulder_id, boulder_name) 
		VALUES ($1, $2) ON CONFLICT (boulder_id) 
		DO UPDATE SET boulder_name = $2`,
		boulder.Id, boulder.Name,
	)
	return err
}

func (d *Postgres) GetBoulder(boulderID string) (*pb.Boulder, error) {
	rows, err := d.db.Query(
		`SELECT 
		b.boulder_id, 
		b.boulder_name,
		a.approach_id,
		a.approach_name,
		a.approach_description,
		a.approach_data
		FROM boulder AS b LEFT JOIN approach AS a 
		ON b.boulder_id = a.boulder_id WHERE b.boulder_id = $1`,
		boulderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	boulder := &pb.Boulder{}
	approaches := []*pb.Approach{}
	for rows.Next() {
		approach := &pb.Approach{}
		approachData := [][]byte{}
		err := rows.Scan(
			&boulder.Id,
			&boulder.Name,
			&approach.Id,
			&approach.Name,
			&approach.Description,
			pq.Array(&approachData))
		if err != nil {
			return nil, err
		}
		for _, data := range approachData {
			point := &pb.Point{}
			err = proto.Unmarshal(data, point)
			if err != nil {
				return nil, err
			}
			approach.Points = append(approach.Points, point)
		}
		approaches = append(approaches, approach)
	}
	boulder.Approaches = approaches

	return boulder, nil
}

func (d *Postgres) SaveApproach(boulderID string, approach *pb.Approach) error {
	// Check if boulder exists
	row := d.db.QueryRow("SELECT boulder_id FROM boulder WHERE boulder_id = $1", boulderID)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return err
	}

	var pointsBytes [][]byte
	for _, point := range approach.Points {
		pointBytes, err := proto.Marshal(point)
		if err != nil {
			return err
		}
		pointsBytes = append(pointsBytes, pointBytes)
	}

	_, err = d.db.Exec(
		`INSERT INTO approach (approach_id, approach_name, approach_description, approach_data, boulder_id) 
		VALUES ($1, $2, $3, $4, $5) ON CONFLICT (approach_id) 
		DO UPDATE SET approach_name = $2, approach_description = $3, approach_data = $4`,
		approach.Id, approach.Name, approach.Description, pq.Array(pointsBytes), boulderID)
	return err
}

func (d *Postgres) GetApproach(approachID string) (*pb.Approach, error) {
	row := d.db.QueryRow(
		`SELECT 
		approach_id, 
		approach_name, 
		approach_description, 
		approach_data 
		FROM approach WHERE approach_id = $1`,
		approachID)

	approach := &pb.Approach{}
	var pointsBytes [][]byte
	err := row.Scan(approach.Id, approach.Name, approach.Description, pq.Array(&pointsBytes))
	if err != nil {
		return nil, err
	}

	for _, pointBytes := range pointsBytes {
		point := &pb.Point{}
		err = proto.Unmarshal(pointBytes, point)
		if err != nil {
			return nil, err
		}
		approach.Points = append(approach.Points, point)
	}

	return approach, nil
}
