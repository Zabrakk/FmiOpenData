package models

type ExplainedParam struct {
	Label          string        `xml:"label"`
	BasePhenomenon string        `xml:"basePhenomenon"`
	UOM            UnitOfMeasure `xml:"uom"`
}

type UnitOfMeasure struct {
	Value string `xml:"uom,attr"`
}
