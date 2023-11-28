// Code generated by avrogen. DO NOT EDIT.

package recordDefault

import (
	"github.com/arenko-group/avro/avrotypegen"
)

type Foo struct {
	F1 int
	F2 string
	F3 string
}

// AvroRecord implements the avro.AvroRecord interface.
func (Foo) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"F1","type":"int"},{"name":"F2","type":"string"},{"default":"hello","name":"F3","type":"string"}],"name":"Foo","type":"record"}`,
		Required: []bool{
			0: true,
			1: true,
		},
		Defaults: []func() interface{}{
			2: func() interface{} {
				return "hello"
			},
		},
	}
}

type R struct {
	RecordField Foo `json:"recordField"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"default":{"F1":44,"F2":"whee","F3":"ok"},"name":"recordField","type":{"fields":[{"name":"F1","type":"int"},{"name":"F2","type":"string"},{"default":"hello","name":"F3","type":"string"}],"name":"Foo","type":"record"}}],"name":"R","type":"record"}`,
		Defaults: []func() interface{}{
			0: func() interface{} {
				return Foo{
					F1: 44,
					F2: "whee",
					F3: "ok",
				}
			},
		},
	}
}
