# FmiOpenData

This is a Go package that allows you to fetch data from the Finnish Meteorological Institute's (FMI) open data web service with various stored queries.

An example workflow for using this package is as follows:
1. Create a stored query object based on what data you want to see, e.g. *q := GetForecastStoredQuery()*
2. Set any query parameters, e.g. *q.SetPlace("Helsinki")*
3. Get the results of that query with *fmiopendata.GetQueryResult(q)*
4. Do what you want with the results

Include this package in your code with:
```golang
import (
	fmiopendata "github.com/Zabrakk/FmiOpenData"
)
```

## Suppored stored queries
The FmiOpenData Go package supports the following stored queries:

### ecmwf::forecast::surface::point::timevaluepair
ECMWF weather forecast fetched to a specific location.
```golang
fmiopendata.GetForecastStoredQuery()
```

### fmi::forecast::silam::airquality::surface::point::timevaluepair
SILAM air quality forecast.
```golang
fmiopendata.GetAirQualityForecastStoredQuery()
```

### fmi::observations::weather::multipointcoverage
Real time weather observations from weather stations.
```golang
fmiopendata.GetRealTimeObservationsStoredQuery()
```

### fmi::observations::weather::daily::timevaluepair
Daily weather observations from weather stations
```golang
fmiopendata.GetDailyObservationsStoredQuery()
```

### fmi::observations::weather::hourly::timevaluepair
Hourly weather observations from weather stations.
```golang
fmiopendata.GetHourlyObservationsStoredQuery()
```

### fmi::forecast::wam::point::timevaluepair
Wave height forecast.
```golang
fmiopendata.GetWaveHeightForecastStoredQuery()
```

More information on the stored queries can be found [here](https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services).

## Setting query parameters
Each stored query has different query parameters associated with it that affect what data the open data web service returns. Here is a list of the parameters that the FmiOpenData package provides a setter for:
```golang
StartTime	 string
EndTime		 string
Parameters	 []string
Timestep	 int
Bbox		 string
LatLon		 string
Place 		 string
Fmisid 		 int
MaxLocations int
```
Note that this package does not support all of the query parameters listed on the [FMI Open data WFS services](https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services) page.

## Explain measurements found in result data
The data provided by the FMI open data web service includes various names for measurements and forecasted values. You can find information on these using two functions provided by this package
```golang
// For observations
fmiopendata.ExplainObservationParam("rrday")
/* Shows the following
Param rrday:
Label:           Precipitation amount
Base Phenomenon: Amount of precipitation
Unit of measure: mm
*/

// For forecasts
fmiopendata.ExplainForecastParam("aqindex")
```

## Code examples
You can find code examples on using this package in the [examples](examples) directory.

## Unit tests
Unit tests can be run with the following command:
```bash
go test ./...
```

## Useful links that helped with creating this package:

https://en.ilmatieteenlaitos.fi/open-data-manual

https://en.ilmatieteenlaitos.fi/open-data-manual-accessing-data

https://en.ilmatieteenlaitos.fi/open-data-manual-wfs-examples-and-guidelines

https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services

http://opendata.fmi.fi/meta?observableProperty=observation

Find fmisids here:
https://en.ilmatieteenlaitos.fi/observation-stations
