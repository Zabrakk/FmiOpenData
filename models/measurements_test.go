package models

import (
	"time"
	"testing"
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

/*
 DAILY
*/
var m1 = MeasurementTVP{
	Time: "2024-06-24T16:00:00Z",
	Value: "10.5",
}
var m2 = MeasurementTVP{
	Time: "2024-06-24T17:00:00Z",
	Value: "11.5",
}

var ms = []Point{
	{ Measurement: m1 },
	{ Measurement: m2 },
}

var dams = AllMeasrurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{ Name: "rrday", Measurements: ms },
		{ Name: "tday", Measurements: ms },
		{ Name: "tmin", Measurements: ms },
		{ Name: "tmax", Measurements: ms },
		{ Name: "TG_PT12H_min", Measurements: ms },
		{ Name: "snow", Measurements: ms },
	},
}

var missing_dams = AllMeasrurements{
	MeasurementTimeseries: []MeasurementTimeseries{},
}

var empty_dams = AllMeasrurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{ Name: "rrday", Measurements: []Point{} },
		{ Name: "tday", Measurements: []Point{} },
		{ Name: "tmin", Measurements: []Point{} },
		{ Name: "tmax", Measurements: []Point{} },
		{ Name: "TG_PT12H_min", Measurements: []Point{} },
		{ Name: "snow", Measurements: []Point{} },
	},
}

func compare_mtvp_slices(a []MeasurementTVP, b []MeasurementTVP) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func check_measurementTVP_slice_is_correct(t *testing.T, result []MeasurementTVP, expected []MeasurementTVP) {
	if !compare_mtvp_slices(result, expected) {
		t.Fatalf("%q != %q", result, expected)
	}
}

func check_measuremenTVP_is_correct(t *testing.T, result MeasurementTVP, expected MeasurementTVP) {
	if result != expected {
		t.Fatalf("%q != %q", result, expected)
	}
}

func TestDailyPrecipitations(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, dams.DailyPrecipitations(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_dams.DailyPrecipitations(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_dams.DailyPrecipitations(), expected)
}

func TestLatestDailyPrecipitation(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, dams.LatestDailyPrecipitation(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_dams.LatestDailyPrecipitation(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_dams.LatestDailyPrecipitation(), m)
}

func TestDailyAirTemperatures(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, dams.DailyAirTemperatures(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_dams.DailyAirTemperatures(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_dams.DailyAirTemperatures(), expected)
}

func TestLatestDailyAirTemperature(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, dams.LatestDailyAirTemperature(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_dams.LatestDailyAirTemperature(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_dams.LatestDailyAirTemperature(), m)
}

func TestDailyMinTemperatures(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, dams.DailyMinTemperatures(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_dams.DailyMinTemperatures(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_dams.DailyMinTemperatures(), expected)
}

func TestLatestDailyMinTemperature(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, dams.LatestDailyMinTemperature(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_dams.LatestDailyMinTemperature(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_dams.LatestDailyMinTemperature(), m)
}

func TestDailyMaxTemperatures(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, dams.DailyMaxTemperatures(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_dams.DailyMaxTemperatures(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_dams.DailyMaxTemperatures(), expected)
}

func TestLatestDailyMaxTemperature(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, dams.LatestDailyMaxTemperature(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_dams.LatestDailyMaxTemperature(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_dams.LatestDailyMaxTemperature(), m)
}

func TestDailyGroundMinTemperatures(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, dams.DailyGroundMinTemperatures(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_dams.DailyGroundMinTemperatures(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_dams.DailyGroundMinTemperatures(), expected)
}

func TestLatestDailyGroundMinTemperature(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, dams.LatestDailyGroundMinTemperature(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_dams.LatestDailyGroundMinTemperature(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_dams.LatestDailyGroundMinTemperature(), m)
}

func TestDailySnowDepths(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, dams.DailySnowDepths(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_dams.DailySnowDepths(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_dams.DailySnowDepths(), expected)
}

func TestLatestDailySnowDepth(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, dams.LatestDailySnowDepth(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_dams.LatestDailySnowDepth(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_dams.LatestDailySnowDepth(), m)
}

// HOURLY

var hams = AllMeasrurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{ Name: "TA_PT1H_AVG", Measurements: ms },
		{ Name: "TA_PT1H_MAX", Measurements: ms},
		{ Name: "TA_PT1H_MIN", Measurements: ms },
		{ Name: "RH_PT1H_AVG", Measurements: ms },
		{ Name: "WS_PT1H_AVG", Measurements: ms },
		{ Name: "WS_PT1H_MAX", Measurements: ms },
		{ Name: "WS_PT1H_MIN", Measurements: ms },
		{ Name: "WD_PT1H_AVG", Measurements: ms },
		{ Name: "PRA_PT1H_ACC", Measurements: ms },
		{ Name: "PRI_PT1H_MAX", Measurements: ms },
		{ Name: "PA_PT1H_AVG", Measurements: ms },
	},
}

var missing_hams = AllMeasrurements{
	MeasurementTimeseries: []MeasurementTimeseries{},
}

var empty_hams = AllMeasrurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{ Name: "TA_PT1H_AVG", Measurements: []Point{} },
		{ Name: "TA_PT1H_MAX", Measurements: []Point{} },
		{ Name: "TA_PT1H_MIN", Measurements: []Point{} },
		{ Name: "RH_PT1H_AVG", Measurements: []Point{} },
		{ Name: "WS_PT1H_AVG", Measurements: []Point{} },
		{ Name: "WS_PT1H_MAX", Measurements: []Point{} },
		{ Name: "WS_PT1H_MIN", Measurements: []Point{} },
		{ Name: "WD_PT1H_AVG", Measurements: []Point{} },
		{ Name: "PRA_PT1H_ACC", Measurements: []Point{} },
		{ Name: "PRI_PT1H_MAX", Measurements: []Point{} },
		{ Name: "PA_PT1H_AVG", Measurements: []Point{} },
	},
}