package http

import (
	"fmt"
	"io"
	nhttp "net/http"
	"github.com/Zabrakk/FmiOpenData/internal/models"

)

func get(url string) string {
	resp, err := nhttp.Get(url)
	if (err != nil) {
		fmt.Println("An error occured")
		fmt.Println(err)
		return ""
	} else {
		fmt.Println("Get was successful")
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return string(body)
	}
}

// TODO Interface

func SendQuery(query models.ObservationQuery) string {
	url := "https://opendata.fmi.fi/wfs?service=WFS&version=2.0.0&request=GetFeature&storedquery_id="
	url += query.Id
	if len(query.Place) == 0 {
		// TODO: Change string?
		url += "&place=" + string(query.Fmisid)
	} else {
		url += "&place=" + query.Place
	}
	fmt.Println(url)
	result := get(url)
	fmt.Println(result)
	return result
}
