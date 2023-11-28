// Code generated by avrogen. DO NOT EDIT.

package largeRecord

import (
	"github.com/arenko-group/avro/avrotypegen"
)

// Common information related to the event which must be included in any clean event

type Data1 struct {
	// Unique identifier for the event used for de-duplication and tracing.
	Uuid *UUID1 `json:"uuid"`

	// Fully qualified name of the host that generated the event that generated the data.
	Hostname *string `json:"hostname"`

	// Trace information not redundant with this object
	Trace *Trace1 `json:"trace"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (Data1) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"doc":"Common information related to the event which must be included in any clean event","fields":[{"default":null,"doc":"Unique identifier for the event used for de-duplication and tracing.","name":"uuid","type":["null",{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"UUID1","namespace":"bodyworks.datatype","type":"record"}]},{"default":null,"doc":"Fully qualified name of the host that generated the event that generated the data.","name":"hostname","type":["null","string"]},{"default":null,"doc":"Trace information not redundant with this object","name":"trace","type":["null",{"doc":"Trace1","fields":[{"default":null,"doc":"Trace Identifier","name":"traceId","type":["null",{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"UUID0","namespace":"headerworks.datatype","type":"record"}]}],"name":"Trace1","type":"record"}]}],"name":"bodyworks.Data1","type":"record"}`,
	}
}

// Trace1

type Trace1 struct {
	// Trace Identifier
	TraceId *UUID0 `json:"traceId"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (Trace1) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"doc":"Trace1","fields":[{"default":null,"doc":"Trace Identifier","name":"traceId","type":["null",{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"UUID0","namespace":"headerworks.datatype","type":"record"}]}],"name":"bodyworks.Trace1","type":"record"}`,
	}
}

// A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014

type UUID1 struct {
	Uuid string `json:"uuid"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (UUID1) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"bodyworks.datatype.UUID1","type":"record"}`,
	}
}

// GoGen test

type Sample struct {
	// Core data information required for any event
	Header *Data0 `json:"header"`

	// Core data information required for any event
	Body *Data1 `json:"body"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (Sample) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"doc":"GoGen test","fields":[{"default":null,"doc":"Core data information required for any event","name":"header","type":["null",{"doc":"Common information related to the event which must be included in any clean event","fields":[{"default":null,"doc":"Unique identifier for the event used for de-duplication and tracing.","name":"uuid","type":["null",{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"UUID0","namespace":"headerworks.datatype","type":"record"}]},{"default":null,"doc":"Fully qualified name of the host that generated the event that generated the data.","name":"hostname","type":["null","string"]},{"default":null,"doc":"Trace information not redundant with this object","name":"trace","type":["null",{"doc":"Trace0","fields":[{"default":null,"doc":"Trace Identifier","name":"traceId","type":["null","headerworks.datatype.UUID0"]}],"name":"Trace0","type":"record"}]}],"name":"Data0","namespace":"headerworks","type":"record"}]},{"default":null,"doc":"Core data information required for any event","name":"body","type":["null",{"doc":"Common information related to the event which must be included in any clean event","fields":[{"default":null,"doc":"Unique identifier for the event used for de-duplication and tracing.","name":"uuid","type":["null",{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"UUID1","namespace":"bodyworks.datatype","type":"record"}]},{"default":null,"doc":"Fully qualified name of the host that generated the event that generated the data.","name":"hostname","type":["null","string"]},{"default":null,"doc":"Trace information not redundant with this object","name":"trace","type":["null",{"doc":"Trace1","fields":[{"default":null,"doc":"Trace Identifier","name":"traceId","type":["null","headerworks.datatype.UUID0"]}],"name":"Trace1","type":"record"}]}],"name":"Data1","namespace":"bodyworks","type":"record"}]}],"name":"com.avro.test.sample","type":"record"}`,
	}
}

// Common information related to the event which must be included in any clean event

type Data0 struct {
	// Unique identifier for the event used for de-duplication and tracing.
	Uuid *UUID0 `json:"uuid"`

	// Fully qualified name of the host that generated the event that generated the data.
	Hostname *string `json:"hostname"`

	// Trace information not redundant with this object
	Trace *Trace0 `json:"trace"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (Data0) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"doc":"Common information related to the event which must be included in any clean event","fields":[{"default":null,"doc":"Unique identifier for the event used for de-duplication and tracing.","name":"uuid","type":["null",{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"UUID0","namespace":"headerworks.datatype","type":"record"}]},{"default":null,"doc":"Fully qualified name of the host that generated the event that generated the data.","name":"hostname","type":["null","string"]},{"default":null,"doc":"Trace information not redundant with this object","name":"trace","type":["null",{"doc":"Trace0","fields":[{"default":null,"doc":"Trace Identifier","name":"traceId","type":["null","headerworks.datatype.UUID0"]}],"name":"Trace0","type":"record"}]}],"name":"headerworks.Data0","type":"record"}`,
	}
}

// Trace0

type Trace0 struct {
	// Trace Identifier
	TraceId *UUID0 `json:"traceId"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (Trace0) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"doc":"Trace0","fields":[{"default":null,"doc":"Trace Identifier","name":"traceId","type":["null",{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"UUID0","namespace":"headerworks.datatype","type":"record"}]}],"name":"headerworks.Trace0","type":"record"}`,
	}
}

// A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014

type UUID0 struct {
	Uuid string `json:"uuid"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (UUID0) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"doc":"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014","fields":[{"default":"","name":"uuid","type":"string"}],"name":"headerworks.datatype.UUID0","type":"record"}`,
	}
}
