// Code generated by avrogen. DO NOT EDIT.

package unionInOut

import "github.com/rogpeppe/avro"

type PrimitiveUnionTestRecord struct {
	UnionField interface{}
}

// AvroRecord implements the avro.AvroRecord interface.
func (PrimitiveUnionTestRecord) AvroRecord() avro.RecordInfo {
	return avro.RecordInfo{
		Schema: `{"fields":[{"default":1234,"name":"UnionField","type":["int","long","float","double","string","boolean","null"]}],"name":"PrimitiveUnionTestRecord","type":"record"}`,
		Fields: []avro.FieldInfo{
			0: {
				Default: func() interface{} {
					return 1234
				},
				Info: &avro.TypeInfo{
					Type: new(interface{}),
					Union: []avro.TypeInfo{{
						Type: new(int),
					}, {
						Type: new(int64),
					}, {
						Type: new(float32),
					}, {
						Type: new(float64),
					}, {
						Type: new(string),
					}, {
						Type: new(bool),
					}, {}},
				},
			},
		},
	}
}

// TODO implement MarshalBinary and UnmarshalBinary methods?