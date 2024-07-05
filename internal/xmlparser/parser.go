package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/Zabrakk/FmiOpenData/models"
)

func ParseQueryResult(respBody io.ReadCloser) (models.AllMeasrurements, error) {
	var result models.AllMeasrurements
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
				measurementTimeseries.Name = startElement.Attr[0].Value
				result.MeasurementTimeseries = append(result.MeasurementTimeseries, measurementTimeseries)
			}
		}
	}
	return result, nil
}

func ParseExplainParamResult(respBody io.ReadCloser) (models.ExplainedParam, error) {
	var result models.ExplainedParam
	val, err := io.ReadAll(respBody)
	if err != nil {
		return models.ExplainedParam{}, err
	}
	if err := xml.Unmarshal(val, &result); err != nil {
		return models.ExplainedParam{}, err
	}
	return result, nil
}
