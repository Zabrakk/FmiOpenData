package models

// THESE ARE SYNOP OBSERVATIONS
// TODO: Instantaneous Weather Observations (real-time) fmi::observations::weather::multipointcoverage

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

func (am AllMeasurements) GetMeasurementTimeseriesNames() []string {
	var measurement_names []string
	for _, mts := range am.MeasurementTimeseries {
		measurement_names = append(measurement_names, mts.Name)
	}
	return measurement_names
}

func (am AllMeasurements) GetMeasurementTimeseriesByName(param string) []MeasurementTVP {
	for _, mts := range am.MeasurementTimeseries {
		if mts.Name == param {
			return mts.Measurements
		}
	}
	return nil
}

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

func (mtvp MeasurementTVP) GetValue() (float64, error) {
	return strconv.ParseFloat(mtvp.Value, 64)
}

func (mtvp MeasurementTVP) GetTime() (time.Time, error) {
	loc, err := time.LoadLocation("Europe/Helsinki")
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation("2006-01-02T15:04:05Z", mtvp.Time, loc)
}
