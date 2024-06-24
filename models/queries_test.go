package models

import (
	"testing"
	"time"
)

var test_time = time.Date(2024, 6, 24, 16, 54, 10, 0, time.Local)
var expected_time_str = "2024-06-24T16:54:10Z"

func TestTimeToQueryForm(t *testing.T) {
	result := timeToQueryFormat(test_time)
	if result != expected_time_str {
		t.Fatalf("%q != %q", result, expected_time_str)
	}
}

func TestSetStartTime(t *testing.T) {
	q := ObservationQuery{}
	q.SetStartTime(test_time)
	if q.StartTime != expected_time_str {
		t.Fatalf("%q != %q", q.StartTime, expected_time_str)
	}
}

func TestSetEndTime(t *testing.T) {
	q := ObservationQuery{}
	q.SetEndTime(test_time)
	if q.EndTime != expected_time_str {
		t.Fatalf("%q != %q", q.EndTime, expected_time_str)
	}
}

func TestSetTimestep(t *testing.T) {
	q := ObservationQuery{}
	expected_val := 15
	q.SetTimestep(expected_val)
	if q.Timestep != expected_val {
		t.Fatalf("%d != %d", q.Timestep, expected_val)
	}
}

func compare_slices(a []string, b []string) bool {
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

func TestSetParameters(t *testing.T) {
	q := ObservationQuery{}
	expected_val := []string{"rrday", "tday"}
	q.SetParameters(expected_val)
	if !compare_slices(q.Parameters, expected_val) {
		t.Fatalf("q.Parameters was incorrect")
	}
}

func TestSetBoundingBox(t *testing.T) {
	q := ObservationQuery{}
	expected_val := "22,64,24,68"
	q.SetBoundingBox(expected_val)
	if q.Bbox != expected_val {
		t.Fatalf("%q != %q", q.Bbox, expected_val)
	}
}

func TestSetPlace(t *testing.T) {
	q := ObservationQuery{}
	expected_val := "Helsinki"
	q.SetPlace(expected_val)
	if q.Place != expected_val {
		t.Fatalf("%q != %q", q.Place, expected_val)
	}
}

func TestSetFmisid(t *testing.T) {
	q := ObservationQuery{}
	expected_val := 123
	q.SetFmisid(expected_val)
	if q.Fmisid != expected_val {
		t.Fatalf("%d != %d", q.Fmisid, expected_val)
	}
}

func TestSetMaxLocations(t *testing.T) {
	q := ObservationQuery{}
	expected_val := 10
	q.SetMaxLocations(expected_val)
	if q.MaxLocations != expected_val {
		t.Fatalf("%d != %d", q.MaxLocations, expected_val)
	}
}
