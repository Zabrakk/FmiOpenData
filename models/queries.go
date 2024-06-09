package models

import(
	"time"
)

type storedQuery struct {
	Id			string
	StartTime	string
	EndTime		string
	Parameters	[]string
}

type ObservationQuery struct {
	storedQuery // Anonymous field
	Place 		string
	Fmisid 		int
	Timestep	int
}

func timeToQueryFormat(t time.Time) string {
	return t.Format("2006-01-02T00:00:00Z")
}

func (q *ObservationQuery) SetPlace(place string) {
	q.Place = place
}

func (q *ObservationQuery) SetFmisid(fmisid int) {
	q.Fmisid = fmisid
}

func (q *ObservationQuery) SetStartTime(startTime time.Time) {
	q.StartTime = timeToQueryFormat(startTime)
}

func (q *ObservationQuery) SetEndTime(endTime time.Time) {
	q.EndTime = timeToQueryFormat(endTime)
}

func (q *ObservationQuery) SetTimestep(timestep int) {
	q.Timestep = timestep
}

func (q *ObservationQuery) SetParameters(parameters []string) {
	q.Parameters = parameters
}
