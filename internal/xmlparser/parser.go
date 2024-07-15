package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/Zabrakk/FmiOpenData/models"
)

// Extracts MeasurementTimeseries from the contents provided by a HTTP GET request.
// Extracted MeasurementTimeseries are added to the AllMeasurements this function returns.
func ParseMeasurementTimeseries(respBody io.ReadCloser) (models.AllMeasurements, error) {
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
				measurementTimeseries.Name = parseMeasurementTimeseriesName(startElement.Attr[0].Value)
				result.MeasurementTimeseries = append(result.MeasurementTimeseries, measurementTimeseries)
			}
		}
	}
	return result, nil
}

func parseMeasurementTimeseriesName(name string) string {
	name = strings.Replace(name, "obs-obs-1-1-", "", 1)
	name = strings.Replace(name, "mts-1-1-", "", 1)
	return name
}

// Extracts BsWfsElements from the contents provided by a HTTP GET request.
// Extracted BsWfsElements are converted to MeasurementTVPs and added to the AllMeasurements struct
// this function returns.
func ParseBsWfsElements(respBody io.ReadCloser) (models.AllMeasurements, error) {
	var result models.AllMeasurements
	var bsWfsElements []models.BsWfsElement
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
				bsWfsElements = append(bsWfsElements, simpleMeasurement)
			}
		}
	}
	for _, entry := range bsWfsElements {
		addBsWfsElementToAllMeasurements(entry, &result)
	}
	return result, nil
}

// Converts BsWfsElements to MeasurementTVPs and adds them a AllMeasurements'
// MeasurementTimeseries based on the measurement's name.
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

// Parses the contents of a HTTP GET request into an ExplainParam struct.
// Returns the ExplainedParam struct.
func ParseExplainParam(val []byte) (models.ExplainedParam, error) {
	var result models.ExplainedParam
	if err := xml.Unmarshal(val, &result); err != nil {
		return models.ExplainedParam{}, err
	}
	return result, nil
}
