package http

import (
	"fmt"
	"io"
	nhttp "net/http"

	"github.com/Zabrakk/FmiOpenData/models"
)

func get(url string) (io.ReadCloser, error) {
	resp, err := nhttp.Get(url)
	if err != nil {
		fmt.Println("An error occured while trying to Get " + url)
		return nil, err
	}
	return resp.Body, nil
}

func GetQueryResult(query models.ObservationQuery) (io.ReadCloser, error) {
	url := query.ToString()
	fmt.Println(url)
	return get(url)
}

func GetUrl(url string) (io.ReadCloser, error) {
	fmt.Println(url)
	return get(url)
}
