package xmlparser

import (
	"testing"
)

func TestParseExplainParameterResult(t *testing.T) {
	xmlFile := `<ObservableProperty>
		<label>Test1</label>
		<basePhenomenon>Test2</basePhenomenon>
		<uom uom="Test3"/>
	</ObservableProperty>`
	result, err := ParseExplainParamResult([]byte(xmlFile))
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

func TestParseExplainParamResultMissingField(t *testing.T) {
	xmlFile := `<ObservableProperty>
		<label>Test1</label>
		<basePhenomenon>Test2</basePhenomenon>
	</ObservableProperty>`
	result, err := ParseExplainParamResult([]byte(xmlFile))
	if err != nil {
		t.Fatalf("%s", err)
	}
	if result.UOM.Value != "" {
		t.Fatalf("Did not handle missing field correctly")
	}
}
