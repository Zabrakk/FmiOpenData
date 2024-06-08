package models

import "encoding/xml"

type MeasurementTimeseries struct {
	MeasurementName string
	XMLName xml.Name `xml:"MeasurementTimeseries"`
	Points []Point `xml:"point"`
}

type Point struct {
	XMLName xml.Name `xml:"point"`
	Observation Obs `xml:"MeasurementTVP"`
}

type Obs struct {
	XMLName xml.Name `xml:"MeasurementTVP"`
	Time string	`xml:"time"`
	Value string `xml:"value"`
}