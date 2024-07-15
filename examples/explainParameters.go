package examples

import (
	"fmt"

	fmiopendata "github.com/Zabrakk/FmiOpenData"
)

/*
  This code example shows how you can use the fmiopendata package to
  find information on the types of measurements present in FMI open data
  observations and forecasts.
*/

func main() {
	fmiopendata.ExplainForecastParam("aqindex")
	fmt.Println()
	p := fmiopendata.ExplainObservationParam("snow")
	fmt.Println("Snow depth is measured in", p.UOM.Value)
}
