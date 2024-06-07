package http

import (
	"fmt"
	"io"
	nhttp "net/http"
)

func Get() {
	resp, err := nhttp.Get("http://opendata.fmi.fi/wfs/eng?request=getFeature&storedquery_id=fmi::ef::stations&networkid=137&")
	if (err != nil) {
		fmt.Println("An error occured")
		fmt.Println(err)
	} else {
		fmt.Println("Get was successful")
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

}