package examples

import (
	"fmt"

	fmiopendata "github.com/Zabrakk/FmiOpenData"
)

/*
  This code example shows how you can fetch a weather forecast for the Helsinki
  Vuosaari harbour (set with Fmisid found at https://en.ilmatieteenlaitos.fi/observation-stations).

*/

func main() {
	// Create stored query for weather forecast
	q := fmiopendata.GetForecastStoredQuery()
	// Set the location with Fmisid
	q.SetFmisid(151028)
	// Get the query's result
	forecast := fmiopendata.GetQueryResult(q)
	// Print out the names of forecasted weather phenomenons
	fmt.Println(forecast.GetMeasurementTimeseriesNames())
	// Get the temperature forecast. Name found with the command above
	temperatures := forecast.GetMeasurementTimeseriesByName("Temperature")
	// Print out one of the entries in the forecast
	fmt.Printf("Temperature will be %s°C at %s\n", temperatures[3].Value, temperatures[3].Time)
	// We can also get a forecast entry's value as float64 and time as a time.Time with the following
	val, err := temperatures[3].GetValue()
	if err != nil {
		fmt.Print("Conversion to float64 failed")
	}
	time, err := temperatures[3].GetTime()
	if err != nil {
		fmt.Println("Conversion to time.Time failed")
	}
	fmt.Println(val, time)
}
