package http

import (
	"io"
	"fmt"
	nhttp "net/http"

	"github.com/Zabrakk/FmiOpenData/internal/models"

)

func get(url string) (io.ReadCloser, error) {
	resp, err := nhttp.Get(url)
	if err != nil {
		fmt.Println("An error occured while trying to Get " + url)
		return nil, err
	}
	return resp.Body, nil
}

// TODO Interfaces

func GetQueryResult(query models.ObservationQuery) (io.ReadCloser, error) {
	url := "https://opendata.fmi.fi/wfs?service=WFS&version=2.0.0&request=GetFeature&storedquery_id="
	url += query.Id
	if len(query.Place) == 0 {
		url += "&fmisid=" + fmt.Sprint(query.Fmisid)
	} else {
		url += "&place=" + query.Place
	}
	fmt.Println(url)
	result, err := get(url)
	return result, err
}
