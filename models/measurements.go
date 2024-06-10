package models

// THESE ARE SYNOP OBSERVATIONS
// TODO: Instantaneous Weather Observations (real-time) fmi::observations::weather::multipointcoverage

import (
	"strings"
	"encoding/xml"
)

type AllMeasrurements struct {
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

func get_mts(am AllMeasrurements, name string) []MeasurementTVP {
	var result []MeasurementTVP
	for _, mts := range am.MeasurementTimeseries {
		if strings.Contains(mts.Name, name) {
			for _, point := range mts.Measurements {
				result = append(result, point.Measurement)
			}
			return result
		}
	}
	return nil
}

/*
 * DAILY OBSERVATIONS
*/

func (am AllMeasrurements) DailyPrecipitations() []MeasurementTVP {
	return get_mts(am, "rrday")
}

func (am AllMeasrurements) DailyAirTemperatures() []MeasurementTVP {
	return get_mts(am, "tday")
}

func (am AllMeasrurements) DailyMinTemperatures() []MeasurementTVP {
	return get_mts(am, "tmin")
}

func (am AllMeasrurements) DailyMaxTemperatures() []MeasurementTVP {
	return get_mts(am, "tmax")
}

func (am AllMeasrurements) DailyGroundMinimumTemperatures() []MeasurementTVP {
	return get_mts(am, "TG_PT12H_min")
}

func (am AllMeasrurements) DailySnowDepths() []MeasurementTVP {
	return get_mts(am, "snow")
}

/*
 * HOURLY OBSERVATIONS
*/

func (am AllMeasrurements) HourlyAirTemperatures() []MeasurementTVP {
	return get_mts(am, "TA_PT1H_AVG")
}

func (am AllMeasrurements) HourlyMaxTemperatures() []MeasurementTVP {
	return get_mts(am, "TA_PT1H_MAX")
}

func (am AllMeasrurements) HourlyMinTemperatures() []MeasurementTVP {
	return get_mts(am, "TA_PT1H_MAX")
}

func (am AllMeasrurements) HourlyRelativehumidities() []MeasurementTVP {
	return get_mts(am, "RH_PT1H_AVG")
}

func (am AllMeasrurements) HourlyWindSpeeds() []MeasurementTVP {
	return get_mts(am, "WS_PT1H_AVG")
}

func (am AllMeasrurements) HourlyMaxWindSpeeds() []MeasurementTVP {
	return get_mts(am, "WS_PT1H_MAX")
}

func (am AllMeasrurements) HourlyMinWindSpeeds() []MeasurementTVP {
	return get_mts(am, "WS_PT1H_MIN")
}

func (am AllMeasrurements) HourlyWindDirections() []MeasurementTVP {
	return get_mts(am, "WD_PT1H_AVG")
}

func (am AllMeasrurements) HourlyPrecipitations() []MeasurementTVP {
	return get_mts(am, "PRA_PT1H_ACC")
}

func (am AllMeasrurements) HourlyMaxPrecipitationIntensities() []MeasurementTVP {
	return get_mts(am, "PRI_PT1H_MAX")
}

func (am AllMeasrurements) HourlyAirPressures() []MeasurementTVP {
	return get_mts(am, "PA_PT1H_AVG")
}