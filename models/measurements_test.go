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

var ms = []MeasurementTVP{m1, m2}

var dams = AllMeasurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{ Name: "rrday", Measurements: ms },
		{ Name: "tday", Measurements: ms },
		{ Name: "tmin", Measurements: ms },
		{ Name: "tmax", Measurements: ms },
		{ Name: "TG_PT12H_min", Measurements: ms },
		{ Name: "snow", Measurements: ms },
	},
}

var missing_dams = AllMeasurements{
	MeasurementTimeseries: []MeasurementTimeseries{},
}

var empty_dams = AllMeasurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{ Name: "rrday", Measurements: []MeasurementTVP{} },
		{ Name: "tday", Measurements: []MeasurementTVP{} },
		{ Name: "tmin", Measurements: []MeasurementTVP{} },
		{ Name: "tmax", Measurements: []MeasurementTVP{} },
		{ Name: "TG_PT12H_min", Measurements: []MeasurementTVP{} },
		{ Name: "snow", Measurements: []MeasurementTVP{} },
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

var hams = AllMeasurements{
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

var missing_hams = AllMeasurements{
	MeasurementTimeseries: []MeasurementTimeseries{},
}

var empty_hams = AllMeasurements{
	MeasurementTimeseries: []MeasurementTimeseries{
		{ Name: "TA_PT1H_AVG", Measurements: []MeasurementTVP{} },
		{ Name: "TA_PT1H_MAX", Measurements: []MeasurementTVP{} },
		{ Name: "TA_PT1H_MIN", Measurements: []MeasurementTVP{} },
		{ Name: "RH_PT1H_AVG", Measurements: []MeasurementTVP{} },
		{ Name: "WS_PT1H_AVG", Measurements: []MeasurementTVP{} },
		{ Name: "WS_PT1H_MAX", Measurements: []MeasurementTVP{} },
		{ Name: "WS_PT1H_MIN", Measurements: []MeasurementTVP{} },
		{ Name: "WD_PT1H_AVG", Measurements: []MeasurementTVP{} },
		{ Name: "PRA_PT1H_ACC", Measurements: []MeasurementTVP{} },
		{ Name: "PRI_PT1H_MAX", Measurements: []MeasurementTVP{} },
		{ Name: "PA_PT1H_AVG", Measurements: []MeasurementTVP{} },
	},
}

func TestHourlyAirTemperatures(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyAirTemperatures(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyAirTemperatures(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyAirTemperatures(), expected)
}

func TestLatestHourlyAirTemperature(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyAirTemperature(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyAirTemperature(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyAirTemperature(), m)
}

func TestHourlyMaxTemperatures(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyMaxTemperatures(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyMaxTemperatures(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyMaxTemperatures(), expected)
}

func TestLatestHourlyMaxTemperature(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyMaxTemperature(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyMaxTemperature(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyMaxTemperature(), m)
}

func TestHourlyMinTemperatures(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyMinTemperatures(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyMinTemperatures(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyMinTemperatures(), expected)
}

func TestLatestHourlyMinTemperature(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyMinTemperature(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyMinTemperature(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyMinTemperature(), m)
}

func TestHourlyRelativehumidities(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyRelativehumidities(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyRelativehumidities(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyRelativehumidities(), expected)
}

func TestLatestHourlyRelativehumidity(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyRelativehumidity(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyRelativehumidity(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyRelativehumidity(), m)
}

func TestHourlyWindSpeeds(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyWindSpeeds(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyWindSpeeds(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyWindSpeeds(), expected)
}

func TestLatestHourlyWindSpeed(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyWindSpeed(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyWindSpeed(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyWindSpeed(), m)
}

func TestHourlyMaxWindSpeeds(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyMaxWindSpeeds(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyMaxWindSpeeds(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyMaxWindSpeeds(), expected)
}

func TestLatestHourlyMaxWindSpeed(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyMaxWindSpeed(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyMaxWindSpeed(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyMaxWindSpeed(), m)
}

func TestHourlyMinWindSpeeds(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyMinWindSpeeds(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyMinWindSpeeds(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyMinWindSpeeds(), expected)
}

func TestLatestHourlyMinWindSpeed(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyMinWindSpeed(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyMinWindSpeed(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyMinWindSpeed(), m)
}

func TestHourlyWindDirections(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyWindDirections(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyWindDirections(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyWindDirections(), expected)
}

func TestLatestHourlyWindDirection(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyWindDirection(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyWindDirection(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyWindDirection(), m)
}

func TestHourlyPrecipitations(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyPrecipitations(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyPrecipitations(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyPrecipitations(), expected)
}

func TestLatestHourlyPrecipitation(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyPrecipitation(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyPrecipitation(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyPrecipitation(), m)
}

func TestHourlyMaxPrecipitationIntensities(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyMaxPrecipitationIntensities(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyMaxPrecipitationIntensities(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyMaxPrecipitationIntensities(), expected)
}

func TestLatestHourlyMaxPrecipitationIntensity(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyMaxPrecipitationIntensity(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyMaxPrecipitationIntensity(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyMaxPrecipitationIntensity(), m)
}


func TestHourlyAirPressures(t *testing.T) {
	// Normal
	expected := []MeasurementTVP{m1, m2}
	check_measurementTVP_slice_is_correct(t, hams.HourlyAirPressures(), expected)
	// Data missing
	expected = []MeasurementTVP{}
	check_measurementTVP_slice_is_correct(t, missing_hams.HourlyAirPressures(), expected)
	// Data empty
	check_measurementTVP_slice_is_correct(t, empty_hams.HourlyAirPressures(), expected)
}

func TestLatestHourlyAirPressure(t *testing.T) {
	// Normal
	check_measuremenTVP_is_correct(t, hams.LatestHourlyAirPressure(), m2)
	// Data missing
	m := MeasurementTVP{}
	check_measuremenTVP_is_correct(t, missing_hams.LatestHourlyAirPressure(), m)
	// Data empty
	check_measuremenTVP_is_correct(t, empty_hams.LatestHourlyAirPressure(), m)
}
