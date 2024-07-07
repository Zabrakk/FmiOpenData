package models

import (
	"testing"
	"time"
)

var measurementTVP = MeasurementTVP{
	Time: "2024-06-24T16:54:10Z",
	Value: "10.5",
}

func TestGetValue(t *testing.T) {
	val, err := measurementTVP.GetValue()
	if err != nil {
		t.Fatalf("GetValue() error occured: %q", err)
	}
	if val != 10.5 {
		t.Fatalf("GetValue() result was incorrect. %f != %f", val, 10.5)
	}
}

func TestGetTime(t *testing.T) {
	val, err := measurementTVP.GetTime()
	if err != nil {
		t.Fatalf("GetTime() error occured: %q", err)
	}
	loc, err := time.LoadLocation("Europe/Helsinki")
	if err != nil {
		t.Fatalf("time.LoadLocation failed")
	}
	expected := time.Date(2024, 6, 24, 16, 54, 10, 0, loc)
	if val.Year() != expected.Year() {
		t.Fatalf("Year was incorrect")
	}
	if val.Month() != expected.Month() {
		t.Fatalf("Month was incorrect")
	}
	if val.Day() != expected.Day() {
		t.Fatalf("Year was incorrect")
	}
	if val.Hour() != expected.Hour() {
		t.Fatalf("Hour was incorrect")
	}
	if val.Minute() != expected.Minute() {
		t.Fatalf("Minute was incorrect")
	}
	if val.Second() != expected.Second() {
		t.Fatalf("Second was incorrect")
	}
}


var allMeasurements = AllMeasurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{
			Name: "rrday", Measurements: []MeasurementTVP{
				{ Time: "2024-06-24T15:54:10Z", Value: "10.5" },
				{ Time: "2024-06-24T16:54:10Z", Value: "12.5" },
			},
		},
		{
			Name: "tday", Measurements: []MeasurementTVP{
				{ Time: "2024-06-24T15:54:10Z", Value: "11.0" },
				{ Time: "2024-06-24T16:54:10Z", Value: "13.0" },
			},
		},
	},
}

func TestGetMeasurementTimeseriesNames(t *testing.T) {
	val := allMeasurements.GetMeasurementTimeseriesNames()
	if len(val) != 2 {
		t.Fatalf("GetMeasurementTimeseriesNames returned incorrect number of names")
	}
	if val[0] != "rrday" || val[1] != "tday" {
		t.Fatalf("GetMeasurementTimeseriesNames returned wrong names")
	}
}

func TestGetMeasurementTimeseriesNamesEmpty(t *testing.T) {
	am := AllMeasurements{}
	if len(am.GetMeasurementTimeseriesNames()) != 0 {
		t.Fatalf("GetMeasurementTimeseriesNames did not work correctly with empty struct")
	}
}

func TestGetMeasurementTimeseriesByName(t *testing.T) {
	val := allMeasurements.GetMeasurementTimeseriesByName("rrday")
	if len(val) != 2 {
		t.Fatalf("GetMeasurementTimeseriesByName returned incorrect number of MeasurementTVPs")
	}
	if v, _ := val[0].GetValue(); v != 10.5 {
		t.Fatalf("GetMeasurementTimeseriesByName returned incorrect measurement")
	}
}

func TestGetMeasurementTimeseriesByNameNotFound(t *testing.T) {
	val := allMeasurements.GetMeasurementTimeseriesByName("test")
	if len(val) != 0 {
		t.Fatalf("GetMeasurementTimeseriesByName did not work correctly with non existent measurement")
	}
}

func TestGetLatestMeasurementByName(t *testing.T) {
	val := allMeasurements.GetLatestMeasurementByName("tday")
	if v, _ := val.GetValue(); v != 13.0 {
		t.Fatalf("GetLatestMeasurementByName returned incorrect measurement")
	}
}

func TestGetLatestMeasurementByNameNotFound(t *testing.T) {
	val := allMeasurements.GetLatestMeasurementByName("ttday")
	if v, _ := val.GetValue(); v != 0.0 {
		t.Fatalf("GetLatestMeasurementByName did not work correctly with non existent measurement")
	}
}
