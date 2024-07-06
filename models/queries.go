package models

import (
	"fmt"
	"strings"
	"time"
)

type storedQuery struct {
	Id			string
	StartTime	string
	EndTime		string
	Parameters	[]string
}

type ObservationQuery struct {
	// Unsupported parameters: crs, geoid, wmo, timezone
	storedQuery
	Timestep	int
	Bbox		string
	Place 		string
	Fmisid 		int
	MaxLocations int
}

func timeToQueryFormat(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}

// Requires year, month, day, hour, minutes and seconds.
func (q *ObservationQuery) SetStartTime(startTime time.Time) {
	q.StartTime = timeToQueryFormat(startTime)
}

// Requires year, month, day, hour, minutes and seconds.
func (q *ObservationQuery) SetEndTime(endTime time.Time) {
	q.EndTime = timeToQueryFormat(endTime)
}

// Timestep is in minutes.
func (q *ObservationQuery) SetTimestep(timestep int) {
	q.Timestep = timestep
}

func (q *ObservationQuery) SetParameters(parameters []string) {
	q.Parameters = parameters
}

// Bbox format is 22,64,24,68. First two numbers are the
// coordinates of the lower left corner, the last two are top right corner
func (q *ObservationQuery) SetBoundingBox(bbox string) {
	q.Bbox = bbox
}

func (q *ObservationQuery) SetPlace(place string) {
	q.Place = place
}

func (q *ObservationQuery) SetFmisid(fmisid int) {
	q.Fmisid = fmisid
}

func (q *ObservationQuery) SetMaxLocations(maxLocations int) {
	q.MaxLocations = maxLocations
}

func (q *storedQuery) ToString() string {
	s := "&storedquery_id=" + q.Id
	if len(q.StartTime) > 0 {
		s += "&starttime=" + q.StartTime
	}
	if len(q.EndTime) > 0 {
		s += "&endtime=" + q.EndTime
	}
	if len(q.Parameters) > 0 {
		s += "&parameters=" + strings.Join(q.Parameters[:], ",")
	}
	return s
}

func (q *ObservationQuery) ToString() string {
	s := "https://opendata.fmi.fi/wfs?service=WFS&version=2.0.0&request=GetFeature"
	s += q.storedQuery.ToString()
	if q.Fmisid > 0 {
		s += "&fmisid=" + fmt.Sprint(q.Fmisid)
	} else if len(q.Place) > 0 {
		s += "&place=" + q.Place
	} else if len(q.Bbox) > 0 {
		s += "&bbox=" + q.Bbox
	}
	if q.Timestep > 0 {
		s += "&timestep=" + fmt.Sprint(q.Timestep)
	}
	if q.MaxLocations > 0 {
		s += "&maxlocations=" + fmt.Sprint(q.MaxLocations)
	}
	return s
}