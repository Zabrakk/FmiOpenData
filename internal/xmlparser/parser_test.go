package xmlparser

import (
	"io"
	"strings"
	"testing"

	"github.com/Zabrakk/FmiOpenData/models"
)

func TestParseExplainParameterResult(t *testing.T) {
	xmlFile := `<ObservableProperty>
		<label>Test1</label>
		<basePhenomenon>Test2</basePhenomenon>
		<uom uom="Test3"/>
	</ObservableProperty>`
	result, err := ParseExplainParam([]byte(xmlFile))
	if err != nil {
		t.Fatalf("%s", err)
	}
	if result.Label != "Test1" {
		t.Fatalf("ExplainedParam Label was incorrect: %s", result.Label)
	}
	if result.BasePhenomenon != "Test2" {
		t.Fatalf("ExplainedParam BasePhenomenon was incorrect: %s", result.BasePhenomenon)
	}
	if result.UOM.Value != "Test3" {
		t.Fatalf("ExplainedParam UOM value was incorrect: %s", result.UOM.Value)
	}
}

func TestParseExplainParamMissingField(t *testing.T) {
	xmlFile := `<ObservableProperty>
		<label>Test1</label>
		<basePhenomenon>Test2</basePhenomenon>
	</ObservableProperty>`
	result, err := ParseExplainParam([]byte(xmlFile))
	if err != nil {
		t.Fatalf("%s", err)
	}
	if result.UOM.Value != "" {
		t.Fatalf("Did not handle missing field correctly")
	}
}

var measurementNames = []string{"rrday", "tday"}
var measurementTimes = []string{"2024-06-09T00:00:00Z", "2024-06-09T06:00:00Z" }
var measurementValues = []string{"16.0", "17.0"}

func TestParseMeasurementTimeseriesName(t *testing.T) {
	name := parseMeasurementTimeseriesName("obs-obs-1-1-Test")
	if name != "Test" {
		t.Fatalf("Incorrect measurement name parsing: %s", name)
	}
	name = parseMeasurementTimeseriesName("mts-1-1-Test")
	if name != "Test" {
		t.Fatalf("Incorrect measurement name parsing: %s", name)
	}
	name = parseMeasurementTimeseriesName("mts-Test2")
	if name != "mts-Test2" {
		t.Fatalf("Incorrect measurement name parsing: %s", name)
	}
}

func TestParseMeasurementTimeseries(t *testing.T) {
	xmlFile := io.NopCloser(strings.NewReader(`
		<wml2:MeasurementTimeseries gml:id="obs-obs-1-1-` + measurementNames[0] + `">
			<wml2:point>
				<wml2:MeasurementTVP>
					<wml2:time>` + measurementTimes[0] + `</wml2:time>
					<wml2:value>` + measurementValues[0] + `</wml2:value>
				</wml2:MeasurementTVP>
			</wml2:point>
			<wml2:point>
				<wml2:MeasurementTVP>
					<wml2:time>` + measurementTimes[1] + `</wml2:time>
					<wml2:value>` + measurementValues[1] + `</wml2:value>
				</wml2:MeasurementTVP>
			</wml2:point>
		</wml2:MeasurementTimeseries>
		<wml2:MeasurementTimeseries gml:id="obs-obs-1-1-` + measurementNames[1] + `">
			<wml2:point>
				<wml2:MeasurementTVP>
					<wml2:time>` + measurementTimes[0] + `</wml2:time>
					<wml2:value>` + measurementValues[0] + `</wml2:value>
				</wml2:MeasurementTVP>
			</wml2:point>
			<wml2:point>
				<wml2:MeasurementTVP>
					<wml2:time>` + measurementTimes[1] + `</wml2:time>
					<wml2:value>` + measurementValues[1] + `</wml2:value>
				</wml2:MeasurementTVP>
			</wml2:point>
		</wml2:MeasurementTimeseries>
	`))
	result, err := ParseMeasurementTimeseries(xmlFile)
	checkAllMeasurementsCorrectness(result, err, t)
}

func TestParseBsWfsElements(t *testing.T) {
	xmlFile := io.NopCloser(strings.NewReader(`
		<wfs:member>
			<BsWfs:BsWfsElement>
				<BsWfs:Time>` + measurementTimes[0] + `</BsWfs:Time>
				<BsWfs:ParameterName>` + measurementNames[0] + `</BsWfs:ParameterName>
				<BsWfs:ParameterValue>` + measurementValues[0] + `</BsWfs:ParameterValue>
			</BsWfs:BsWfsElement>
		</wfs:member>
		<wfs:member>
			<BsWfs:BsWfsElement>
				<BsWfs:Time>` + measurementTimes[0] + `</BsWfs:Time>
				<BsWfs:ParameterName>` + measurementNames[1] + `</BsWfs:ParameterName>
				<BsWfs:ParameterValue>` + measurementValues[0] + `</BsWfs:ParameterValue>
			</BsWfs:BsWfsElement>
		</wfs:member>
		<wfs:member>
			<BsWfs:BsWfsElement>
				<BsWfs:Time>` + measurementTimes[1] + `</BsWfs:Time>
				<BsWfs:ParameterName>` + measurementNames[0] + `</BsWfs:ParameterName>
				<BsWfs:ParameterValue>` + measurementValues[1] + `</BsWfs:ParameterValue>
			</BsWfs:BsWfsElement>
		</wfs:member>
		<wfs:member>
			<BsWfs:BsWfsElement>
				<BsWfs:Time>` + measurementTimes[1] + `</BsWfs:Time>
				<BsWfs:ParameterName>` + measurementNames[1] + `</BsWfs:ParameterName>
				<BsWfs:ParameterValue>` + measurementValues[1] + `</BsWfs:ParameterValue>
			</BsWfs:BsWfsElement>
		</wfs:member>
	`))
	result, err := ParseBsWfsElements(xmlFile)
	checkAllMeasurementsCorrectness(result, err, t)
}

func checkAllMeasurementsCorrectness(result models.AllMeasurements, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("%s", err)
	}
	if len(result.MeasurementTimeseries) != 2 {
		t.Fatalf("Incorrect number of MeasurementTimeseries: %d", len(result.MeasurementTimeseries))
	}
	for i := 0; i < 2; i++ {
		if result.MeasurementTimeseries[i].Name != measurementNames[i] {
			t.Fatalf("Inccorect measurement name for first entry: %s", result.MeasurementTimeseries[i].Name)
		}
		for j:= 0; j < 2; j++ {
			if result.MeasurementTimeseries[i].Measurements[j].Time != measurementTimes[j] {
				t.Fatalf("Incorrect measurement time for entry i:%d j:%d", i, j)
			}
			if result.MeasurementTimeseries[i].Measurements[j].Value != measurementValues[j] {
				t.Fatalf("Incorrect measurement value for entry i:%d j:%d", i, j)
			}
		}
	}
}
