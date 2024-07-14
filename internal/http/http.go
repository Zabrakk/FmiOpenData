package http

import (
	"fmt"
	"io"
	nhttp "net/http"
)

func get(url string) (io.ReadCloser, error) {
	resp, err := nhttp.Get(url)
	if err != nil {
		fmt.Println("An error occured while trying to Get " + url)
		return nil, err
	}
	return resp.Body, nil
}

// Sends a GET requst to the given url. The resulting contents
// is then returned along with any possible errors.
func GetFromUrl(url string) (io.ReadCloser, error) {
	fmt.Println(url)
	return get(url)
}
