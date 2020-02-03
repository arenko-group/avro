// Code generated by avrogen. DO NOT EDIT.

package duplicateRecord

import (
	"github.com/heetch/avro/avrotypegen"
)

type R1 struct {
	F R2
	G R2
	H int
}

// AvroRecord implements the avro.AvroRecord interface.
func (R1) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"F","type":{"fields":[{"name":"A","type":"string"}],"name":"R2","type":"record"}},{"name":"G","type":"R2"},{"name":"H","type":"int"}],"name":"R1","type":"record"}`,
		Required: []bool{
			0: true,
			1: true,
			2: true,
		},
	}
}

type R2 struct {
	A string
}

// AvroRecord implements the avro.AvroRecord interface.
func (R2) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"A","type":"string"}],"name":"R2","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}
