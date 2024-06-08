package models

type StoredQuery struct {
	Id			string
	StartTime	string
	EndTime		string
	Parameters	[]string
}

type ObservationQuery struct {
	StoredQuery // Anonymous field
	Place 		string
	Fmisid 		int
	timestep	int
}
