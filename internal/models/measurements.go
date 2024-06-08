package models

import (
	"strings"
	"encoding/xml"
)

type DailyMeasrurements struct {
	MeasurementTimeseries []MeasurementTimeseries
}

// Measurement timeseries includes measurement name, e.g. obs-obs-1-1-tday,
// and an array of measurement points.
type MeasurementTimeseries struct {
	Name string
	XMLName xml.Name `xml:"MeasurementTimeseries"`
	Measurements []Point `xml:"point"`
}

// A measurement point includes one measurement time-value pair
type Point struct {
	XMLName xml.Name `xml:"point"`
	Measurement MeasurementTVP `xml:"MeasurementTVP"`
}

// Measurement Time-Value Pair. Includes the measurement's time and value
type MeasurementTVP struct {
	XMLName xml.Name `xml:"MeasurementTVP"`
	Time string	`xml:"time"`
	Value string `xml:"value"`
}

func get_mts(dm DailyMeasrurements, name string) []MeasurementTVP {
	var result []MeasurementTVP
	for _, mts := range dm.MeasurementTimeseries {
		if strings.Contains(mts.Name, name) {
			for _, point := range mts.Measurements {
				result = append(result, point.Measurement)
			}
			return result
		}
	}
	return nil
}

func (dm DailyMeasrurements) Precipitation() []MeasurementTVP {
	return get_mts(dm, "rrday")
}

func (dm DailyMeasrurements) AirTemperature() []MeasurementTVP {
	return get_mts(dm, "tday")
}

func (dm DailyMeasrurements) MinimumTemperature() []MeasurementTVP {
	return get_mts(dm, "tmin")
}

func (dm DailyMeasrurements) MaximumTemperature() []MeasurementTVP {
	return get_mts(dm, "tmax")
}

func (dm DailyMeasrurements) GroundMinimumTemperature() []MeasurementTVP {
	return get_mts(dm, "TG_PT12H_min")
}

func (dm DailyMeasrurements) SnowDepth() []MeasurementTVP {
	return get_mts(dm, "snow")
}
