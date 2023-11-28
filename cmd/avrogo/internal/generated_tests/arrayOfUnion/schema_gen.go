// Code generated by avrogen. DO NOT EDIT.

package arrayOfUnion

import (
	"github.com/arenko-group/avro/avrotypegen"
)

type R struct {
	// Allowed types for interface{} value:
	// 	int
	// 	string
	F []interface{}
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"F","type":{"items":["int","string"],"type":"array"}}],"name":"R","type":"record"}`,
		Required: []bool{
			0: true,
		},
		Unions: []avrotypegen.UnionInfo{
			0: {
				Type: new([]interface{}),
				Union: []avrotypegen.UnionInfo{{
					Type: new(int),
				}, {
					Type: new(string),
				}},
			},
		},
	}
}
