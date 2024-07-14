package models

import (
	"fmt"
	"strings"
	"time"
)

type Query interface {
	ToString() string
}

type StoredQuery struct {
	Id			string
	StartTime	string
	EndTime		string
	Parameters	[]string
	Timestep	int
	Bbox		string
	LatLon		string
	Place 		string
	Fmisid 		int
	MaxLocations int
}

func timeToQueryFormat(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}

// Requires year, month, day, hour, minutes and seconds.
func (q *StoredQuery) SetStartTime(startTime time.Time) {
	q.StartTime = timeToQueryFormat(startTime)
}

// Requires year, month, day, hour, minutes and seconds.
func (q *StoredQuery) SetEndTime(endTime time.Time) {
	q.EndTime = timeToQueryFormat(endTime)
}

// Timestep is in minutes.
func (q *StoredQuery) SetTimestep(timestep int) {
	q.Timestep = timestep
}

func (q *StoredQuery) SetParameters(parameters []string) {
	q.Parameters = parameters
}

// Bbox format is 22,64,24,68. First two numbers are the
// coordinates of the lower left corner, the last two are top right corner
func (q *StoredQuery) SetBoundingBox(bbox string) {
	q.Bbox = bbox
}

// LatLon format is 60.11,19.90.
func (q *StoredQuery) SetLatLon(latlon string) {
	q.LatLon = latlon
}

func (q *StoredQuery) SetPlace(place string) {
	q.Place = place
}

func (q *StoredQuery) SetFmisid(fmisid int) {
	q.Fmisid = fmisid
}

func (q *StoredQuery) SetMaxLocations(maxLocations int) {
	q.MaxLocations = maxLocations
}

// Returns a string which is a URL created based on the StoredQuery
// struct's field values. The URL can then be used in GETing FMI open data
func (q *StoredQuery) ToString() string {
	s := "https://opendata.fmi.fi/wfs?service=WFS&version=2.0.0&request=GetFeature"
	s += "&storedquery_id=" + q.Id
	if len(q.StartTime) > 0 {
		s += "&starttime=" + q.StartTime
	}
	if len(q.EndTime) > 0 {
		s += "&endtime=" + q.EndTime
	}
	if len(q.Parameters) > 0 {
		s += "&parameters=" + strings.Join(q.Parameters[:], ",")
	}
	if q.Fmisid > 0 {
		s += "&fmisid=" + fmt.Sprint(q.Fmisid)
	}
	if len(q.Place) > 0 {
		s += "&place=" + q.Place
	}
	if len(q.Bbox) > 0 {
		s += "&bbox=" + q.Bbox
	}
	if len(q.LatLon) > 0 {
		s += "&latlon=" + q.LatLon
	}
	if q.Timestep > 0 {
		s += "&timestep=" + fmt.Sprint(q.Timestep)
	}
	if q.MaxLocations > 0 {
		s += "&maxlocations=" + fmt.Sprint(q.MaxLocations)
	}
	return s
}