# FmiOpenData


Useful links:

https://en.ilmatieteenlaitos.fi/open-data-manual

https://en.ilmatieteenlaitos.fi/open-data-manual-accessing-data

https://en.ilmatieteenlaitos.fi/open-data-manual-wfs-examples-and-guidelines

https://en.ilmatieteenlaitos.fi/open-data-manual-fmi-wfs-services

http://opendata.fmi.fi/meta?observableProperty=observation


## Suppored stored queries
The FmiOpenData Go package supports the following stored queries:

```golang
"fmi::observations::weather::daily::timevaluepair"
// Accessed in code with
q := fmiopendata.GetDailyObservationQuery()
```

```golang
"fmi::observations::weather::hourly::timevaluepair"
// Accessed in code with
q := fmiopendata.GetHourlyObservationQuery()
```

```golang
"fmi::observations::weather::multipointcoverage"
// Accessed in code with
q := fmiopendata.GetRealTimeObservationQuery()
```

## Explain measurements found in result data

```golang
fmiopendata.ExplainParam("rrday")
/* Shows the following
Param rrday:
Label:           Precipitation amount
Base Phenomenon: Amount of precipitation
Unit of measure: mm
*/
```
