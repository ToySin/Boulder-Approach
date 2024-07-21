package service

import (
	"encoding/xml"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Tracks  []Track  `xml:"trk"`
}

type Track struct {
	Segments []TrackSegment `xml:trkseg`
}

type TrackSegment struct {
	Points []TrackPoint `xml:trkpt`
}

type TrackPoint struct {
	Latitude  float64 `xml:lat,attr`
	Longitude float64 `xml:lon,attr`
}

func ParseGPX(data []byte) (*GPX, error) {
	gpx := &GPX{}
	err := xml.Unmarshal(data, gpx)
	if err != nil {
		return nil, err
	}
	return gpx, nil
}
