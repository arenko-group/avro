// Code generated by avrogen. DO NOT EDIT.

package recordWithDefaultNotPresent

import "github.com/rogpeppe/avro/avrotypegen"

type R struct {
	A int
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"A","type":"int"}],"name":"R","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}

// TODO implement MarshalBinary and UnmarshalBinary methods?