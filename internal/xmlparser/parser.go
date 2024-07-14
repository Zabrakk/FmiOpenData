package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/Zabrakk/FmiOpenData/models"
)

func ParseQueryResult(respBody io.ReadCloser) (models.AllMeasurements, error) {
	var result models.AllMeasurements
	decoder := xml.NewDecoder(respBody)
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error occured while readig XML token")
			return result, err
		}
		if startElement, ok := token.(xml.StartElement); ok {
			if startElement.Name.Local == "MeasurementTimeseries" {
				var measurementTimeseries models.MeasurementTimeseries
				decoder.DecodeElement(&measurementTimeseries, &startElement)
				measurementTimeseries.Name = strings.Replace(startElement.Attr[0].Value, "obs-obs-1-1-", "", 1)
				result.MeasurementTimeseries = append(result.MeasurementTimeseries, measurementTimeseries)
			}
		}
	}
	return result, nil
}

func ParseSimpleQueryResult(respBody io.ReadCloser) (models.AllMeasurements, error) {
	var result models.AllMeasurements
	var simpleMeasurements []models.BsWfsElement
	decoder := xml.NewDecoder(respBody)
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error occured while readin XML token")
			return result, err
		}
		if startElement, ok := token.(xml.StartElement); ok {
			if startElement.Name.Local == "BsWfsElement" {
				var simpleMeasurement models.BsWfsElement
				decoder.DecodeElement(&simpleMeasurement, &startElement)
				simpleMeasurements = append(simpleMeasurements, simpleMeasurement)
			}
		}
	}
	for _, entry := range simpleMeasurements {
		addBsWfsElementToAllMeasurements(entry, &result)
	}
	return result, nil
}

func addBsWfsElementToAllMeasurements(simple models.BsWfsElement, all *models.AllMeasurements) {
	mtvp := models.MeasurementTVP{Time: simple.Time, Value: simple.ParameterValue}
	measurementTypeIncluded := false
	i := 0
	for idx, entry := range all.MeasurementTimeseries {
		if entry.Name == simple.ParameterName {
			measurementTypeIncluded = true
			i = idx
			break
		}
	}
	if !measurementTypeIncluded {
		mts := models.MeasurementTimeseries{Name: simple.ParameterName, Measurements: []models.MeasurementTVP{mtvp}}
		all.MeasurementTimeseries = append(all.MeasurementTimeseries, mts)
	} else {
		all.MeasurementTimeseries[i].Measurements = append(all.MeasurementTimeseries[i].Measurements, mtvp)
	}
}

func ParseExplainParamResult(val []byte) (models.ExplainedParam, error) {
	var result models.ExplainedParam
	if err := xml.Unmarshal(val, &result); err != nil {
		return models.ExplainedParam{}, err
	}
	return result, nil
}
