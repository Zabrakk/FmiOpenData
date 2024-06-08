package http

import (
	"fmt"
	"io"
	"encoding/xml"
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
		//body, _ := io.ReadAll(resp.Body)
		//return string(body)

		decoder := xml.NewDecoder(resp.Body)
		//fmt.Println(decoder)
		for {
			t, err := decoder.Token()
			if err != nil {
				if err == io.EOF {
					fmt.Println("End of File reached")
					break
				}
				fmt.Println("An error occured")
				break
			}
			//fmt.Println(t)
			switch t := t.(type) {
			case xml.StartElement:
				//fmt.Println(t.Name.Local)
				if t.Name.Local == "MeasurementTVP" {
					//fmt.Println(t.Attr)
					var obs models.Obs
					decoder.DecodeElement(&obs, &t)
					fmt.Println(obs)
					attr := t.Attr
					for _, a := range attr {
						fmt.Println(a.Name)
						fmt.Println(a.Value)
					}
				}
				/*if t.Name.Local == "MeasurementTVP" {
					var obs models.Obs
					if err := decoder.DecodeElement(&obs, &t); err != nil {
						fmt.Println(err.Error())
						return ""
					}
					fmt.Println(obs)
				}*/
			}

		}

		return ""
	}
}

// TODO Interface

func SendQuery(query models.ObservationQuery) string {
	url := "https://opendata.fmi.fi/wfs?service=WFS&version=2.0.0&request=GetFeature&storedquery_id="
	url += query.Id
	if len(query.Place) == 0 {
		// TODO: Change string?
		url += "&fmisid=" + fmt.Sprint(query.Fmisid)
	} else {
		url += "&place=" + query.Place
	}
	fmt.Println(url)
	result := get(url)
	//fmt.Println(result)
	return result
}
