package models

type StoredQuery struct {
	Id string
}

type ObservationQuery struct {
	StoredQuery // Anonymous field
	Place string
	Fmisid int
}

type Observation struct {
	PrecipitationRate float64
	SnowDepth	float64
	TemperatureMean float64
	TemperatureMax	float64
	TemperatureMin	float64
}

