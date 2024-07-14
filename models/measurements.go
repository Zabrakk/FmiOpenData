package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type AllMeasurements struct {
	MeasurementTimeseries []MeasurementTimeseries
}

// Measurement timeseries includes measurement name, e.g. obs-obs-1-1-tday,
// and an array of measurement points.
type MeasurementTimeseries struct {
	Name string
	XMLName xml.Name `xml:"MeasurementTimeseries"`
	Measurements []MeasurementTVP `xml:"point>MeasurementTVP"`
}

// Measurement Time-Value Pair. Includes the measurement's time and value
type MeasurementTVP struct {
	XMLName xml.Name `xml:"MeasurementTVP"`
	Time string	`xml:"time"`
	Value string `xml:"value"`
}

type BsWfsElement struct {
	XMLName xml.Name `xml:"BsWfsElement"`
	Time string	`xml:"Time"`
	ParameterName string `xml:"ParameterName"`
	ParameterValue string `xml:"ParameterValue"`
}

/*
 * AllMeasurements
*/

// Returns the names of measurements in the AllMeasurements struct
func (am AllMeasurements) GetMeasurementTimeseriesNames() []string {
	var measurement_names []string
	for _, mts := range am.MeasurementTimeseries {
		measurement_names = append(measurement_names, mts.Name)
	}
	return measurement_names
}

// Returns all MeasurementTVPs for the given measurement.
// For example, GetMeasurementTimeseriesByName("t2m") would return
// all air temperature measurements in the AllMeasurements.
func (am AllMeasurements) GetMeasurementTimeseriesByName(param string) []MeasurementTVP {
	for _, mts := range am.MeasurementTimeseries {
		if mts.Name == param {
			return mts.Measurements
		}
	}
	return nil
}

// Returns the most recent MeasurementTVP for the given measurement.
// For example, GetLatestMeasurementByName("t2m") would return the
// latest air temperature measurement.
func (am AllMeasurements) GetLatestMeasurementByName(param string) MeasurementTVP {
	mts := am.GetMeasurementTimeseriesByName(param)
	if len(mts) > 0 {
		return mts[len(mts)-1]
	}
	fmt.Printf("No latest measurement found for %s\n", param)
	return MeasurementTVP{}
}

/*
 * MeasurementTVP
*/

// Returns the MeasurementTVP's value as a float64.
func (mtvp MeasurementTVP) GetValue() (float64, error) {
	return strconv.ParseFloat(mtvp.Value, 64)
}

// Returns the time a MeasurementTVP was measured.
func (mtvp MeasurementTVP) GetTime() (time.Time, error) {
	loc, err := time.LoadLocation("Europe/Helsinki")
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation("2006-01-02T15:04:05Z", mtvp.Time, loc)
}
