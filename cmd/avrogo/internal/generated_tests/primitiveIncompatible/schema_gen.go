// Code generated by avrogen. DO NOT EDIT.

package primitiveIncompatible

import (
	"github.com/arenko-group/avro/avrotypegen"
)

type R struct {
	F string `json:"f"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"f","type":"string"}],"name":"R","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}
