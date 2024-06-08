package http

import (
	"io"
	"fmt"
	"strings"
	"errors"
	nhttp "net/http"

	"github.com/Zabrakk/FmiOpenData/internal/models"
)

func queryToUrlString(query models.ObservationQuery) (string, error) {
	if len(query.Place) < 1 && query.Fmisid == 0 {
		return "", errors.New("ERROR: You must specify either Place or Fmisid for the ObservationQuery")
	}
	url := "https://opendata.fmi.fi/wfs?service=WFS&version=2.0.0&request=GetFeature&storedquery_id="
	url += query.Id
	if query.Fmisid > 0 {
		url += "&fmisid=" + fmt.Sprint(query.Fmisid)
	} else {
		url += "&place=" + query.Place
	}
	fmt.Println(query.StartTime)
	if len(query.StartTime) > 0 {
		url += "&starttime=" + query.StartTime
	}
	if len(query.EndTime) > 0 {
		url += "&endtime=" + query.EndTime
	}
	if query.Timestep > 0 {
		url += "&timestep=" + fmt.Sprint(query.Timestep)
	}
	if len(query.Parameters) > 0 {
		url += "&parameters=" + strings.Join(query.Parameters[:], ",")
	}
	return url, nil
}

func get(url string) (io.ReadCloser, error) {
	resp, err := nhttp.Get(url)
	if err != nil {
		fmt.Println("An error occured while trying to Get " + url)
		return nil, err
	}
	return resp.Body, nil
}

func GetQueryResult(query models.ObservationQuery) (io.ReadCloser, error) {
	url, err := queryToUrlString(query)
	if err != nil {
		return nil, err
	}
	fmt.Println(url)
	result, err := get(url)
	return result, err
}
