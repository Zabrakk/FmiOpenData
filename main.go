package fmiopendata

import (
	"fmt"

	"github.com/Zabrakk/FmiOpenData/internal/models"
	"github.com/Zabrakk/FmiOpenData/internal/http"
	"github.com/Zabrakk/FmiOpenData/internal/xmlparser"
)


func GetQueryResult() {
	query := models.GetDailyObservationStructForPlace("Helsinki")
	//query := models.GetDailyObservationStructForFmisid(100968)
	queryResult, err := http.GetQueryResult(query)
	if err != nil {
		fmt.Println("ERROR!!!")
		return
	}
	result, err := xmlparser.ParseQueryResult(queryResult)
	if err != nil {
		fmt.Println("ERROR WHILE PARSING XML!!!")
		return
	}
	for _, obs := range result {
		fmt.Println(obs)
		fmt.Println()
	}
}
