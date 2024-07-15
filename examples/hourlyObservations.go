package examples

import (
	"fmt"
	"time"

	fmiopendata "github.com/Zabrakk/FmiOpenData"
)

/*
  This example code demonstrates how you can query the FMI open data for hourly weather observations
  measured in Helsinki, Finland during the time period 2024/06/05 - 2024/06/08. By setting the "parameters"
  query param's value to ["tday", "rrday"] provided measurements will been narrowed down to air temperatue (tday) and precipitation amount (rrday).
*/

func main() {
	// Create a stored query corresponding to hourly observations
	q := fmiopendata.GetHourlyObservationsStoredQuery()
	// Set query parameters
	q.SetPlace("Helsinki")
	q.SetStartTime(time.Date(2024, 6, 5, 0, 0, 0, 0, time.Local))
	q.SetEndTime(time.Date(2024, 6, 8, 0, 0, 0, 0, time.Local))
	q.SetParameters([]string{"tday", "rrday"})
	// Fetch the data FMI provides for the created stored query
	observations := fmiopendata.GetQueryResult(q)
	// Print out the value of the most recent air temperature measurement from the given time period
	val, _ := observations.GetLatestMeasurementByName("tday").GetValue()
	fmt.Println("Air temperature measurement's value was", val)
}
