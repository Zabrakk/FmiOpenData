package fmiopendata

import (
	"fmt"
	"io"

	"github.com/Zabrakk/FmiOpenData/internal/http"
	"github.com/Zabrakk/FmiOpenData/internal/xmlparser"
	"github.com/Zabrakk/FmiOpenData/models"
)

func GetDailyObservationQuery() models.ObservationQuery {
	query := models.ObservationQuery{}
	query.Id = "fmi::observations::weather::daily::timevaluepair"
	return query
}

func GetHourlyObservationQuery() models.ObservationQuery {
	query := models.ObservationQuery{}
	query.Id = "fmi::observations::weather::hourly::timevaluepair"
	return query
}

func GetQueryResult(query models.ObservationQuery) models.AllMeasurements {
	queryResult, err := http.GetFromUrl(query.ToString())
	if err != nil {
		fmt.Println(err)
		return models.AllMeasurements{}
	}
	defer queryResult.Close() // Close when this function's execution ends.
	result, err := xmlparser.ParseQueryResult(queryResult)
	if err != nil {
		fmt.Println("ERROR WHILE PARSING XML!!!")
		return models.AllMeasurements{}
	}
	return result
}

func ExplainParam(param string) models.ExplainedParam {
	url := "http://opendata.fmi.fi/meta?observableProperty=observation&param=" + param
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
	explainedParam, err := xmlparser.ParseExplainParamResult([]byte(val))
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
