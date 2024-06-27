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

func get_latest(mtvps []MeasurementTVP) MeasurementTVP {
	if len(mtvps) < 1 {
		return MeasurementTVP{}
	}
	return mtvps[len(mtvps)-1]
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
	return get_latest(am.DailyPrecipitations())
}

func (am AllMeasrurements) DailyAirTemperatures() []MeasurementTVP {
	return get_measurements(am, "tday")
}

func (am AllMeasrurements) LatestDailyAirTemperature() MeasurementTVP {
	return get_latest(am.DailyAirTemperatures())
}

func (am AllMeasrurements) DailyMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "tmin")
}

func (am AllMeasrurements) LatestDailyMinTemperature() MeasurementTVP {
	return get_latest(am.DailyMinTemperatures())
}

func (am AllMeasrurements) DailyMaxTemperatures() []MeasurementTVP {
	return get_measurements(am, "tmax")
}

func (am AllMeasrurements) LatestDailyMaxTemperature() MeasurementTVP {
	return get_latest(am.DailyMaxTemperatures())
}

func (am AllMeasrurements) DailyGroundMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "TG_PT12H_min")
}

func (am AllMeasrurements) LatestDailyGroundMinTemperature() MeasurementTVP {
	return get_latest(am.DailyGroundMinTemperatures())
}

func (am AllMeasrurements) DailySnowDepths() []MeasurementTVP {
	return get_measurements(am, "snow")
}

func (am AllMeasrurements) LatestDailySnowDepth() MeasurementTVP {
	return get_latest(am.DailySnowDepths())
}

/*
 * HOURLY OBSERVATIONS
*/

func (am AllMeasrurements) HourlyAirTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyAirTemperature() MeasurementTVP {
	return get_latest(am.HourlyAirTemperatures())
}

func (am AllMeasrurements) HourlyMaxTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_MAX")
}

func (am AllMeasrurements) LatestHourlyMaxTemperature() MeasurementTVP {
	return get_latest(am.HourlyMaxTemperatures())
}

func (am AllMeasrurements) HourlyMinTemperatures() []MeasurementTVP {
	return get_measurements(am, "TA_PT1H_MIN")
}

func (am AllMeasrurements) LatestHourlyMinTemperature() MeasurementTVP {
	return get_latest(am.HourlyMinTemperatures())
}

func (am AllMeasrurements) HourlyRelativehumidities() []MeasurementTVP {
	return get_measurements(am, "RH_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyRelativehumidity() MeasurementTVP {
	return get_latest(am.HourlyRelativehumidities())
}

func (am AllMeasrurements) HourlyWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyWindSpeed() MeasurementTVP {
	return get_latest(am.HourlyWindSpeeds())
}

func (am AllMeasrurements) HourlyMaxWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_MAX")
}

func (am AllMeasrurements) LatestHourlyMaxWindSpeed() MeasurementTVP {
	return get_latest(am.HourlyMaxWindSpeeds())
}

func (am AllMeasrurements) HourlyMinWindSpeeds() []MeasurementTVP {
	return get_measurements(am, "WS_PT1H_MIN")
}

func (am AllMeasrurements) LatestHourlyMinWindSpeed() MeasurementTVP {
	return get_latest(am.HourlyMinWindSpeeds())
}

func (am AllMeasrurements) HourlyWindDirections() []MeasurementTVP {
	return get_measurements(am, "WD_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyWindDirection() MeasurementTVP {
	return get_latest(am.HourlyWindDirections())
}

func (am AllMeasrurements) HourlyPrecipitations() []MeasurementTVP {
	return get_measurements(am, "PRA_PT1H_ACC")
}

func (am AllMeasrurements) LatestHourlyPrecipitation() MeasurementTVP {
	return get_latest(am.HourlyPrecipitations())
}

func (am AllMeasrurements) HourlyMaxPrecipitationIntensities() []MeasurementTVP {
	return get_measurements(am, "PRI_PT1H_MAX")
}

func (am AllMeasrurements) LatestHourlyMaxPrecipitationIntensity() MeasurementTVP {
	return get_latest(am.HourlyMaxPrecipitationIntensities())
}

func (am AllMeasrurements) HourlyAirPressures() []MeasurementTVP {
	return get_measurements(am, "PA_PT1H_AVG")
}

func (am AllMeasrurements) LatestHourlyAirPressure() MeasurementTVP {
	return get_latest(am.HourlyAirPressures())
}
