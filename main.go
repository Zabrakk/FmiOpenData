package fmiopendata

import (
	"fmt"
	"io"
	"strings"

	"github.com/Zabrakk/FmiOpenData/internal/http"
	"github.com/Zabrakk/FmiOpenData/internal/xmlparser"
	"github.com/Zabrakk/FmiOpenData/models"
)

func getStoredQuery(id string) models.StoredQuery {
	return models.StoredQuery{Id: id}
}

// Returns a stored query for "fmi::observations::weather::daily::timevaluepair"
// Supported query parameters are:
//  starttime
//  endtime
//  timestep
//  parameters
//  bbox
//  place
//  fmisid
//  maxlocations
// Check https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services for more info.
func GetDailyObservationsStoredQuery() models.StoredQuery {
	return getStoredQuery("fmi::observations::weather::daily::timevaluepair")
}

// Returns a stored query for "fmi::observations::weather::hourly::timevaluepair"
// Supported query parameters are:
//  starttime
//  endtime
//  timestep
//  parameters
//  bbox
//  place
//  fmisid
//  maxlocations
// Check https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services for more info.
func GetHourlyObservationsStoredQuery() models.StoredQuery {
	return getStoredQuery("fmi::observations::weather::hourly::timevaluepair")
}

// Returns a stored query for "fmi::observations::weather::simple"
// Supported query parameters are:
//  starttime
//  endtime
//  timestep
//  parameters
//  bbox
//  place
//  fmisid
//  maxlocations
// Check https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services for more info.
func GetRealTimeObservationsStoredQuery() models.StoredQuery {
	return getStoredQuery("fmi::observations::weather::simple")
}

// Returns a stored query for "ecmwf::forecast::surface::point::timevaluepair".
// Supported query parameters are:
//  starttime
//  endtime
//  timestep
//  parameters
//  place
//  latlon
//  fmisid
// Check https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services for more info.
func GetForecastStoredQuery() models.StoredQuery {
	return getStoredQuery("ecmwf::forecast::surface::point::timevaluepair")
}

// Returns a stored query for "fmi::forecast::silam::airquality::surface::point::timevaluepair".
// Supported query parameters are:
//  starttime
//  endtime
//  timestep
//  parameters
//  latlon
//  fmisid
// Check https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services for more info.
func GetAirQualityForecast() models.StoredQuery {
	return getStoredQuery("fmi::forecast::silam::airquality::surface::point::timevaluepair")
}

// Returns a stored query for "fmi::forecast::wam::point::timevaluepair".
// Supported query parameters are:
//  starttime
//  endtime
//  timestep
//  parameters
//  fmisid
//  latlon
//  timezone
// Check https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services for more info.
func GetWaveHeightForecast() models.StoredQuery {
	return getStoredQuery("fmi::forecast::wam::point::timevaluepair")
}

// Performs a GET request, based on the given StoredQuery, to FMI.
// Returns the data provided by the GET request in a AllMeasurements struct.
func GetQueryResult(query models.StoredQuery) models.AllMeasurements {
	queryResult, err := http.GetFromUrl(query.ToString())
	if err != nil {
		fmt.Println(err)
		return models.AllMeasurements{}
	}
	defer queryResult.Close() // Close when this function's execution ends.
	var result models.AllMeasurements
	if strings.Contains(query.Id, "simple") {
		result, err = xmlparser.ParseBsWfsElements(queryResult)
	} else {
		result, err = xmlparser.ParseMeasurementTimeseries(queryResult)
	}
	if err != nil {
		fmt.Println("ERROR WHILE PARSING XML!!!")
		return models.AllMeasurements{}
	}
	return result
}

// This function prints out the information FMI provides on a given observation
// measurement parameter, e.g. t2m. The info is also returned in an ExplainedParam struct.
// Forecast parameters are not supported.
func ExplainObservationParam(param string) models.ExplainedParam {
	return explainParamFromUrl("http://opendata.fmi.fi/meta?observableProperty=observation", param)
}

// This function prints out the information FMI provides on a given forecast
// measurement parameter, e.g., aqindex. This info is also returned in an ExplainedParam struct.
// Observation parameters are not supported.
func ExplainForecastParam(param string) models.ExplainedParam {
	return explainParamFromUrl("http://opendata.fmi.fi/meta?observableProperty=forecast", param)
}

func explainParamFromUrl(url string, param string) models.ExplainedParam {
	url += "&param=" + param
	result, err := http.GetFromUrl(url)
	if err != nil {
		fmt.Println(err)
		return models.ExplainedParam{}
	}
	defer result.Close()
	val, err := io.ReadAll(result)
	if err != nil {
		fmt.Println(err)
		return models.ExplainedParam{}
	}
	explainedParam, err := xmlparser.ParseExplainParam([]byte(val))
	if err != nil {
		fmt.Println(err)
		return models.ExplainedParam{}
	}
	fmt.Println("Param " + param + ":")
	fmt.Println("Label:           " + explainedParam.Label)
	fmt.Println("Base Phenomenon: " + explainedParam.BasePhenomenon)
	fmt.Println("Unit of measure: " + explainedParam.UOM.Value)
	return explainedParam
}
