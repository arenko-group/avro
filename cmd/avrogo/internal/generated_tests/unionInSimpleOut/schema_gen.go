// Code generated by avrogen. DO NOT EDIT.

package unionInSimpleOut

import (
	"github.com/arenko-group/avro/avrotypegen"
)

type R struct {
	UnionField string
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"UnionField","type":"string"}],"name":"R","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}
