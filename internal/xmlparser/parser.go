package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/Zabrakk/FmiOpenData/internal/models"
)

func ParseQueryResult(respBody io.ReadCloser) ([]models.MeasurementTimeseries, error) {
	var result []models.MeasurementTimeseries
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
				measurementTimeseries.MeasurementName = startElement.Attr[0].Value
				result = append(result, measurementTimeseries)
			}
		}
	}
	return result, nil
}

