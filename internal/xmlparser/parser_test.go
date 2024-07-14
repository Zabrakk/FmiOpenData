package xmlparser

import (
	"io"
	"strings"
	"testing"
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

func TestParseMeasurementTimeseries(t *testing.T) {
	xmlFile := strings.NewReader(`
		<wml2:MeasurementTimeseries gml:id="obs-obs-1-1-rrday">
			<wml2:point>
				<wml2:MeasurementTVP>
					<wml2:time>2024-06-09T00:00:00Z</wml2:time>
					<wml2:value>16.0</wml2:value>
				</wml2:MeasurementTVP>
			</wml2:point>
			<wml2:point>
				<wml2:MeasurementTVP>
					<wml2:time>2024-06-09T06:00:00Z</wml2:time>
					<wml2:value>17.0</wml2:value>
				</wml2:MeasurementTVP>
			</wml2:point>
		</wml2:MeasurementTimeseries>
		<wml2:MeasurementTimeseries gml:id="obs-obs-1-1-tday">
			<wml2:point>
				<wml2:MeasurementTVP>
					<wml2:time>2024-06-09T00:00:00Z</wml2:time>
					<wml2:value>16.0</wml2:value>
				</wml2:MeasurementTVP>
			</wml2:point>
			<wml2:point>
				<wml2:MeasurementTVP>
					<wml2:time>2024-06-09T06:00:00Z</wml2:time>
					<wml2:value>17.0</wml2:value>
				</wml2:MeasurementTVP>
			</wml2:point>
		</wml2:MeasurementTimeseries>
	`)
	xmlFile2 := io.NopCloser(xmlFile)
	result, err := ParseMeasurementTimeseries(xmlFile2)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if len(result.MeasurementTimeseries) != 2 {
		t.Fatalf("Incorrect number of MeasurementTimeseries: %d", len(result.MeasurementTimeseries))
	}
	measurementNames := []string{"rrday", "tday"}
	measurementTimes := []string{"2024-06-09T00:00:00Z", "2024-06-09T06:00:00Z", }
	measurementValues := []string{"16.0", "17.0"}
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

func TestParseBsWfsElements(t *testing.T) {
	xmlFile := io.NopCloser(strings.NewReader(`
		<wfs:member>
			<BsWfs:BsWfsElement>
				<BsWfs:Time>2024-06-09T00:00:00Z</BsWfs:Time>
				<BsWfs:ParameterName>t2m</BsWfs:ParameterName>
				<BsWfs:ParameterValue>16.0</BsWfs:ParameterValue>
			</BsWfs:BsWfsElement>
		</wfs:member>
		<wfs:member>
			<BsWfs:BsWfsElement>
				<BsWfs:Time>2024-06-09T00:00:00Z</BsWfs:Time>
				<BsWfs:ParameterName>ws_10min</BsWfs:ParameterName>
				<BsWfs:ParameterValue>16.0</BsWfs:ParameterValue>
			</BsWfs:BsWfsElement>
		</wfs:member>
		<wfs:member>
			<BsWfs:BsWfsElement>
				<BsWfs:Time>2024-06-09T06:00:00Z</BsWfs:Time>
				<BsWfs:ParameterName>t2m</BsWfs:ParameterName>
				<BsWfs:ParameterValue>17.0</BsWfs:ParameterValue>
			</BsWfs:BsWfsElement>
		</wfs:member>
		<wfs:member>
			<BsWfs:BsWfsElement>
				<BsWfs:Time>2024-06-09T06:00:00Z</BsWfs:Time>
				<BsWfs:ParameterName>ws_10min</BsWfs:ParameterName>
				<BsWfs:ParameterValue>17.0</BsWfs:ParameterValue>
			</BsWfs:BsWfsElement>
		</wfs:member>
	`))
	result, err := ParseBsWfsElements(xmlFile)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if len(result.MeasurementTimeseries) != 2 {
		t.Fatalf("Incorrect number of MeasurementTimeseries: %d", len(result.MeasurementTimeseries))
	}
	measurementNames := []string{"t2m", "ws_10min"}
	measurementTimes := []string{"2024-06-09T00:00:00Z", "2024-06-09T06:00:00Z" }
	measurementValues := []string{"16.0", "17.0"}
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