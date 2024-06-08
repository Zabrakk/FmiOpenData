package models

//import "encoding/xml"

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

type Observation struct {
	PrecipitationRate float64
	SnowDepth	float64
	TemperatureMean float64
	TemperatureMax	float64
	TemperatureMin	float64
}


type Obs struct {
	//XMLName xml.Name `xml:"MeasurementTVP"`
	Time string	`xml:"time"`
	Value string `xml:"value"`
}