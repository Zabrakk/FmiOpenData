package models

// THESE ARE SYNOP OBSERVATIONS
// TODO: Instantaneous Weather Observations (real-time) fmi::observations::weather::multipointcoverage

import (
	"time"
	"strconv"
	"strings"
	"encoding/xml"
)

type AllMeasrurements struct {
	MeasurementTimeseries []MeasurementTimeseries
}

// Measurement timeseries includes measurement name, e.g. obs-obs-1-1-tday,
// and an array of measurement points.
type MeasurementTimeseries struct {
	Name string
	XMLName xml.Name `xml:"MeasurementTimeseries"`
	Measurements []Point `xml:"point"`
}

// A measurement point includes one measurement time-value pair
type Point struct {
	XMLName xml.Name `xml:"point"`
	Measurement MeasurementTVP `xml:"MeasurementTVP"`
}

// Measurement Time-Value Pair. Includes the measurement's time and value
type MeasurementTVP struct {
	XMLName xml.Name `xml:"MeasurementTVP"`
	Time string	`xml:"time"`
	Value string `xml:"value"`
}

func get_measurements(am AllMeasrurements, name string) []MeasurementTVP {
	var result []MeasurementTVP
	for _, mts := range am.MeasurementTimeseries {
		if strings.Contains(mts.Name, name) {
			for _, point := range mts.Measurements {
				result = append(result, point.Measurement)
			}
			return result
		}
	}
	return nil
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

func (am AllMeasrurements) DailyPrecipitations() []MeasurementTVP {
	return get_measurements(am, "rrday")
}

func (am AllMeasrurements) LatestDailyPrecipitation() MeasurementTVP {
	mtvps := am.DailyPrecipitations()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) DailyAirTemperatures() []MeasurementTVP {
	return get_measurements(am, "tday")
}

func (am AllMeasrurements) LatestDailyAirTemperature() MeasurementTVP {
	mtvps := am.DailyAirTemperatures()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) DailyMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "tmin")
}

func (am AllMeasrurements) LatestDailyMinTemperature() MeasurementTVP {
	mtvps := am.DailyMinTemperatures()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) DailyMaxTemperatures() []MeasurementTVP {
	return get_measurements(am, "tmax")
}

func (am AllMeasrurements) LatestDailyMaxTemperature() MeasurementTVP {
	mtvps := am.DailyMaxTemperatures()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) DailyGroundMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "TG_PT12H_min")
}

func (am AllMeasrurements) LatestDailyGroundMinTemperature() MeasurementTVP {
	mtvps := am.DailyGroundMinTemperatures()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) DailySnowDepths() []MeasurementTVP {
	return get_measurements(am, "snow")
}

func (am AllMeasrurements) LatestDailySnowDepth() MeasurementTVP {
	mtvps := am.DailySnowDepths()
	return mtvps[len(mtvps)-1]
}

/*
 * HOURLY OBSERVATIONS
*/

func (am AllMeasrurements) HourlyAirTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyAirTemperature() MeasurementTVP {
	mtvps := am.HourlyAirTemperatures()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyMaxTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_MAX")
}

func (am AllMeasrurements) LatestHourlyMaxTemperature() MeasurementTVP {
	mtvps := am.HourlyMaxTemperatures()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_MAX")
}

func (am AllMeasrurements) LatestHourlyMinTemperature() MeasurementTVP {
	mtvps := am.HourlyMinTemperatures()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyRelativehumidities() []MeasurementTVP {
	return get_measurements(am, "RH_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyRelativehumidity() MeasurementTVP {
	mtvps := am.HourlyRelativehumidities()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyWindSpeed() MeasurementTVP {
	mtvps := am.HourlyWindSpeeds()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyMaxWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_MAX")
}

func (am AllMeasrurements) LatestHourlyMaxWindSpeed() MeasurementTVP {
	mtvps := am.HourlyMaxWindSpeeds()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyMinWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_MIN")
}

func (am AllMeasrurements) LatestHourlyMinWindSpeed() MeasurementTVP {
	mtvps := am.HourlyMinWindSpeeds()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyWindDirections() []MeasurementTVP {
	return get_measurements(am, "WD_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyWindDirection() MeasurementTVP {
	mtvps := am.HourlyWindDirections()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyPrecipitations() []MeasurementTVP {
	return get_measurements(am, "PRA_PT1H_ACC")
}

func (am AllMeasrurements) LatestHourlyPrecipitation() MeasurementTVP {
	mtvps := am.HourlyPrecipitations()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyMaxPrecipitationIntensities() []MeasurementTVP {
	return get_measurements(am, "PRI_PT1H_MAX")
}

func (am AllMeasrurements) LatestHourlyMaxPrecipitationIntensity() MeasurementTVP {
	mtvps := am.HourlyMaxPrecipitationIntensities()
	return mtvps[len(mtvps)-1]
}

func (am AllMeasrurements) HourlyAirPressures() []MeasurementTVP {
	return get_measurements(am, "PA_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyAirPressure() MeasurementTVP {
	mtvps := am.HourlyAirPressures()
	return mtvps[len(mtvps)-1]
}
