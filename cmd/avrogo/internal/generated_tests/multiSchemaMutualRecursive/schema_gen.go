// Code generated by avrogen. DO NOT EDIT.

package multiSchemaMutualRecursive

import (
	"github.com/arenko-group/avro/avrotypegen"
)

type R struct {
	F []S
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"default":[],"name":"F","type":{"items":{"fields":[{"default":"","name":"Data","type":"string"},{"default":null,"name":"Child","type":["null","R"]}],"name":"S","type":"record"},"type":"array"}}],"name":"R","type":"record"}`,
	}
}
