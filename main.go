package fmiopendata

import (
	"fmt"

	"github.com/Zabrakk/FmiOpenData/internal/models"
	"github.com/Zabrakk/FmiOpenData/internal/http"
	"github.com/Zabrakk/FmiOpenData/internal/xmlparser"
)


func GetQueryResult() {
	query := models.GetDailyObservationStructForPlace("Helsinki")
	queryResult, err := http.GetQueryResult(query)
	defer queryResult.Close() // Close when this function's execution ends.
	if err != nil {
		fmt.Println("ERROR!!!")
		return
	}
	result, err := xmlparser.ParseQueryResult(queryResult)
	if err != nil {
		fmt.Println("ERROR WHILE PARSING XML!!!")
		return
	}
	for _, m := range result.Precipitation() {
		fmt.Println(m.Value, m.Time)
	}
}
