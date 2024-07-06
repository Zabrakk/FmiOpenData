package models

// THESE ARE SYNOP OBSERVATIONS
// TODO: Instantaneous Weather Observations (real-time) fmi::observations::weather::multipointcoverage

import (
	"encoding/xml"
	"strconv"
	"strings"
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

func get_measurements(am AllMeasurements, name string) []MeasurementTVP {
	for _, mts := range am.MeasurementTimeseries {
		if strings.Contains(mts.Name, name) {
			return mts.Measurements
		}
	}
	return nil
}

func get_latest(mtvps []MeasurementTVP) MeasurementTVP {
	if len(mtvps) < 1 {
		return MeasurementTVP{}
	}
	return mtvps[len(mtvps)-1]
}

func (am AllMeasurements) GetMeasurementNames() []string {
	var measurement_names []string
	for _, mts := range am.MeasurementTimeseries {
		measurement_names = append(measurement_names, mts.Name)
	}
	return measurement_names
}

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

/*
 * DAILY OBSERVATIONS
*/

func (am AllMeasurements) DailyPrecipitations() []MeasurementTVP {
	return get_measurements(am, "rrday")
}

func (am AllMeasurements) LatestDailyPrecipitation() MeasurementTVP {
	return get_latest(am.DailyPrecipitations())
}

func (am AllMeasurements) DailyAirTemperatures() []MeasurementTVP {
	return get_measurements(am, "tday")
}

func (am AllMeasurements) LatestDailyAirTemperature() MeasurementTVP {
	return get_latest(am.DailyAirTemperatures())
}

func (am AllMeasurements) DailyMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "tmin")
}

func (am AllMeasurements) LatestDailyMinTemperature() MeasurementTVP {
	return get_latest(am.DailyMinTemperatures())
}

func (am AllMeasurements) DailyMaxTemperatures() []MeasurementTVP {
	return get_measurements(am, "tmax")
}

func (am AllMeasurements) LatestDailyMaxTemperature() MeasurementTVP {
	return get_latest(am.DailyMaxTemperatures())
}

func (am AllMeasurements) DailyGroundMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "TG_PT12H_min")
}

func (am AllMeasurements) LatestDailyGroundMinTemperature() MeasurementTVP {
	return get_latest(am.DailyGroundMinTemperatures())
}

func (am AllMeasurements) DailySnowDepths() []MeasurementTVP {
	return get_measurements(am, "snow")
}

func (am AllMeasurements) LatestDailySnowDepth() MeasurementTVP {
	return get_latest(am.DailySnowDepths())
}

/*
 * HOURLY OBSERVATIONS
*/

func (am AllMeasurements) HourlyAirTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_AVG")
}

func (am AllMeasurements) LatestHourlyAirTemperature() MeasurementTVP {
	return get_latest(am.HourlyAirTemperatures())
}

func (am AllMeasurements) HourlyMaxTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_MAX")
}

func (am AllMeasurements) LatestHourlyMaxTemperature() MeasurementTVP {
	return get_latest(am.HourlyMaxTemperatures())
}

func (am AllMeasurements) HourlyMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_MIN")
}

func (am AllMeasurements) LatestHourlyMinTemperature() MeasurementTVP {
	return get_latest(am.HourlyMinTemperatures())
}

func (am AllMeasurements) HourlyRelativehumidities() []MeasurementTVP {
	return get_measurements(am, "RH_PT1H_AVG")
}

func (am AllMeasurements) LatestHourlyRelativehumidity() MeasurementTVP {
	return get_latest(am.HourlyRelativehumidities())
}

func (am AllMeasurements) HourlyWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_AVG")
}

func (am AllMeasurements) LatestHourlyWindSpeed() MeasurementTVP {
	return get_latest(am.HourlyWindSpeeds())
}

func (am AllMeasurements) HourlyMaxWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_MAX")
}

func (am AllMeasurements) LatestHourlyMaxWindSpeed() MeasurementTVP {
	return get_latest(am.HourlyMaxWindSpeeds())
}

func (am AllMeasurements) HourlyMinWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_MIN")
}

func (am AllMeasurements) LatestHourlyMinWindSpeed() MeasurementTVP {
	return get_latest(am.HourlyMinWindSpeeds())
}

func (am AllMeasurements) HourlyWindDirections() []MeasurementTVP {
	return get_measurements(am, "WD_PT1H_AVG")
}

func (am AllMeasurements) LatestHourlyWindDirection() MeasurementTVP {
	return get_latest(am.HourlyWindDirections())
}

func (am AllMeasurements) HourlyPrecipitations() []MeasurementTVP {
	return get_measurements(am, "PRA_PT1H_ACC")
}

func (am AllMeasurements) LatestHourlyPrecipitation() MeasurementTVP {
	return get_latest(am.HourlyPrecipitations())
}

func (am AllMeasurements) HourlyMaxPrecipitationIntensities() []MeasurementTVP {
	return get_measurements(am, "PRI_PT1H_MAX")
}

func (am AllMeasurements) LatestHourlyMaxPrecipitationIntensity() MeasurementTVP {
	return get_latest(am.HourlyMaxPrecipitationIntensities())
}

func (am AllMeasurements) HourlyAirPressures() []MeasurementTVP {
	return get_measurements(am, "PA_PT1H_AVG")
}

func (am AllMeasurements) LatestHourlyAirPressure() MeasurementTVP {
	return get_latest(am.HourlyAirPressures())
}
