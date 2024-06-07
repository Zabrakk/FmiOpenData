package fmiopendata

import (
	"fmt"
	"github.com/Zabrakk/FmiOpenData/internal/models"
	"github.com/Zabrakk/FmiOpenData/internal/http"
)


func Get() {
	http.Get()
}

/*
func GetModel() {
	query := models.ObservationQuery{models.StoredQuery{"fmi::observation"}, "Helsinki", 12345}
	fmt.Println(query.Id, query.Place, query.Fmisid)
}
*/

func GetQuery() {
	fmt.Println(models.GetDailyObservationStructForPlace("Helsinki"))
	fmt.Println(models.GetDailyObservationStructForFmisid(12345))
}

func PrintText(text string) {
	fmt.Print(text)
}
