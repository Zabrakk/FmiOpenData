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
		{
			Name: "rrday",
			Measurements: ms,
		},
		{
			Name: "tday",
			Measurements: ms,
		},
		{
			Name: "tmin",
			Measurements: ms,
		},
		{
			Name: "tmax",
			Measurements: ms,
		},
		{
			Name: "TG_PT12H_min",
			Measurements: ms,
		},
		{
			Name: "snow",
			Measurements: ms,
		},
	},
}

var missing_dams = AllMeasrurements{
	MeasurementTimeseries: []MeasurementTimeseries{},
}

var empty_dams = AllMeasrurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{
			Name: "rrday",
			Measurements: []Point{},
		},
		{
			Name: "tday",
			Measurements: []Point{},
		},
		{
			Name: "tmin",
			Measurements: []Point{},
		},
		{
			Name: "tmax",
			Measurements: []Point{},
		},
		{
			Name: "TG_PT12H_min",
			Measurements: []Point{},
		},
		{
			Name: "snow",
			Measurements: []Point{},
		},
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

func check_measurementTVP_slice_is_correct(t *testing.T, m1 []MeasurementTVP, m2 []MeasurementTVP) {
	if !compare_mtvp_slices(m1, m2) {
		t.Fatalf("%q != %q", m1, m2)
	}
}

func check_measuremenTVP_is_correct(t *testing.T, m1 MeasurementTVP, m2 MeasurementTVP) {
	if m1 != m2 {
		t.Fatalf("%q != %q", m1, m2)
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
	expected := []MeasurementTVP{m1, m2}
	r := dams.DailyAirTemperatures()
	if !compare_mtvp_slices(r, expected) {
		t.Fatalf("%q != %q", r, expected)
	}
}

func TestLatestDailyAirTemperature(t *testing.T) {
	r := dams.LatestDailyAirTemperature()
	if r != m2 {
		t.Fatalf("%q != %q", r, m2)
	}
}

func TestDailyMinTemperatures(t *testing.T) {
	expected := []MeasurementTVP{m1, m2}
	r := dams.DailyMinTemperatures()
	if !compare_mtvp_slices(r, expected) {
		t.Fatalf("%q != %q", r, expected)
	}
}

func TestLatestDailyMinTemperature(t *testing.T) {
	r := dams.LatestDailyMinTemperature()
	if r != m2 {
		t.Fatalf("%q != %q", r, m2)
	}
}

func TestDailyMaxTemperatures(t *testing.T) {
	expected := []MeasurementTVP{m1, m2}
	r := dams.DailyMaxTemperatures()
	if !compare_mtvp_slices(r, expected) {
		t.Fatalf("%q != %q", r, expected)
	}
}

func TestLatestDailyMaxTemperature(t *testing.T) {
	r := dams.LatestDailyMaxTemperature()
	if r != m2 {
		t.Fatalf("%q != %q", r, m2)
	}
}

func TestDailyGroundMinTemperatures(t *testing.T) {
	expected := []MeasurementTVP{m1, m2}
	r := dams.DailyGroundMinTemperatures()
	if !compare_mtvp_slices(r, expected) {
		t.Fatalf("%q != %q", r, expected)
	}
}

func TestLatestDailyGroundMinTemperature(t *testing.T) {
	r := dams.LatestDailyGroundMinTemperature()
	if r != m2 {
		t.Fatalf("%q != %q", r, m2)
	}
}

func TestDailySnowDepths(t *testing.T) {
	expected := []MeasurementTVP{m1, m2}
	r := dams.DailySnowDepths()
	if !compare_mtvp_slices(r, expected) {
		t.Fatalf("%q != %q", r, expected)
	}
}

func TestLatestDailySnowDepth(t *testing.T) {
	r := dams.LatestDailySnowDepth()
	if r != m2 {
		t.Fatalf("%q != %q", r, m2)
	}
}
