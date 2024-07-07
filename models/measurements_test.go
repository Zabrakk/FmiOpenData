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
