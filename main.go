package fmiopendata

import (
	"fmt"
	"time"

	"github.com/Zabrakk/FmiOpenData/internal/models"
	"github.com/Zabrakk/FmiOpenData/internal/http"
	"github.com/Zabrakk/FmiOpenData/internal/xmlparser"
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

func TimeToQueryFormat(t time.Time) string {
	return t.Format("2006-01-02T00:00:00Z")
}

func GetQueryResult(query models.ObservationQuery) {
	queryResult, err := http.GetQueryResult(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer queryResult.Close() // Close when this function's execution ends.
	result, err := xmlparser.ParseQueryResult(queryResult)
	if err != nil {
		fmt.Println("ERROR WHILE PARSING XML!!!")
		return
	}
	for _, m := range result.Precipitation() {
		fmt.Println(m.Value, m.Time)
	}
}