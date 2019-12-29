package avro

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"sort"

	"github.com/actgardner/gogen-avro/schema"
)

// Set to true for deterministic output.
const sortMapKeys = false

// Marshal encodes the given value as
// a message using the Avro binary encoding.
//
// See https://avro.apache.org/docs/current/spec.html#binary_encoding
//
// Currently, x must be a type that was generated
// by the avro-generate-go command.
func Marshal(x interface{}) (_ []byte, marshalErr error) {
	var recInfo RecordInfo

	if x, ok := x.(AvroRecord); ok {
		recInfo = x.AvroRecord()
	} else {
		// TODO generate schema from type
		return nil, fmt.Errorf("cannot get schema info for %T", x)
	}
	namespace := schema.NewNamespace(false)
	at, err := namespace.TypeForSchema([]byte(recInfo.Schema))
	if err != nil {
		return nil, fmt.Errorf("cannot parse schema; %v", err)
	}
	xv := reflect.ValueOf(x)
	info, err := newAzTypeInfo(xv.Type())
	if err != nil {
		return nil, err
	}
	enc := typeEncoder(namespace, at, xv.Type(), info)
	var e encodeState
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(*encodeError); ok {
				marshalErr = err.err
			} else {
				panic(r)
			}
		}
	}()
	enc(&e, xv)
	return e.Bytes(), nil
}

type encodeState struct {
	bytes.Buffer
	scratch [64]byte
}

// error aborts the encoding by panicking with err wrapped in encodeError.
func (e *encodeState) error(err error) {
	panic(&encodeError{err})
}

func errorEncoder(err error) encoderFunc {
	return func(e *encodeState, v reflect.Value) {
		e.error(err)
	}
}

type encodeError struct {
	err error
}

type encoderFunc func(e *encodeState, v reflect.Value)

func typeEncoder(ns *schema.Namespace, at schema.AvroType, t reflect.Type, info azTypeInfo) encoderFunc {
	// TODO cache this so it's faster and so that we can deal with recursive types.
	switch at := at.(type) {
	case *schema.Reference:
		def, ok := ns.Definitions[at.TypeName]
		if !ok {
			return errorEncoder(fmt.Errorf("reference to unknown name %q", at.TypeName))
		}
		switch def := def.(type) {
		case *schema.RecordDefinition:
			if t.Kind() != reflect.Struct {
				return errorEncoder(fmt.Errorf("expected struct"))
			}
			if len(info.entries) == 0 {
				// The type itself might contribute information.
				info1, err := newAzTypeInfo(info.ftype)
				if err != nil {
					return errorEncoder(fmt.Errorf("cannot get info for %s: %v", info.ftype, err))
				}
				info = info1
			}
			if len(info.entries) != len(def.Fields()) {
				return errorEncoder(fmt.Errorf("entry count mismatch (info entries %d vs definition fields %d; %s vs %s)", len(info.entries), len(def.Fields()), t, def.Name()))
			}
			if t.NumField() != len(def.Fields()) {
				return errorEncoder(fmt.Errorf("field count mismatch (%d vs %d; %s vs %s)", t.NumField(), len(def.Fields()), t, def.Name()))
			}
			fields := make([]encoderFunc, len(def.Fields()))
			for i, f := range def.Fields() {
				fields[i] = typeEncoder(ns, f.Type(), t.Field(i).Type, info.entries[i])
			}
			return structEncoder{fields}.encode
		case *schema.EnumDefinition:
			return longEncoder
		case *schema.FixedDefinition:
			return fixedEncoder{def.SizeBytes()}.encode
		default:
			return errorEncoder(fmt.Errorf("unknown definition type %T", def))
		}
	case *schema.UnionField:
		atypes := at.ItemTypes()
		switch t.Kind() {
		case reflect.Ptr:
			// It's a union of null and one other type, represented by a Go pointer.
			if len(atypes) != 2 {
				return errorEncoder(fmt.Errorf("unexpected item type count in union"))
			}
			switch {
			case info.entries[0].ftype == nil:
				return ptrUnionEncoder{
					indexes:    [2]byte{0, 1},
					encodeElem: typeEncoder(ns, atypes[1], info.entries[1].ftype, info.entries[1]),
				}.encode
			case info.entries[1].ftype == nil:
				return ptrUnionEncoder{
					indexes:    [2]byte{1, 0},
					encodeElem: typeEncoder(ns, atypes[0], info.entries[0].ftype, info.entries[0]),
				}.encode
			default:
				return errorEncoder(fmt.Errorf("unexpected types in union"))
			}
		case reflect.Interface:
			enc := unionEncoder{
				nullIndex: -1,
				choices:   make([]unionEncoderChoice, len(info.entries)),
			}
			for i, entry := range info.entries {
				if entry.ftype == nil {
					enc.nullIndex = i
				} else {
					enc.choices[i] = unionEncoderChoice{
						typ: entry.ftype,
						enc: typeEncoder(ns, atypes[i], entry.ftype, info.entries[i]),
					}
				}
			}
			return enc.encode
		default:
			return errorEncoder(fmt.Errorf("union type is not pointer or interface"))
		}
	case *schema.MapField:
		return mapEncoder{typeEncoder(ns, at.ItemType(), t.Elem(), info)}.encode
	case *schema.ArrayField:
		return arrayEncoder{typeEncoder(ns, at.ItemType(), t.Elem(), info)}.encode
	case *schema.BoolField:
		return boolEncoder
	case *schema.IntField,
		*schema.LongField:
		return longEncoder
	case *schema.FloatField:
		return floatEncoder
	case *schema.DoubleField:
		return doubleEncoder
	case *schema.BytesField:
		return bytesEncoder
	case *schema.StringField:
		return stringEncoder
	default:
		return errorEncoder(fmt.Errorf("unknown avro schema type %T", at))
	}
}

type fixedEncoder struct {
	size int
}

func (fe fixedEncoder) encode(e *encodeState, v reflect.Value) {
	if v.CanAddr() {
		e.Write(v.Slice(0, fe.size).Bytes())
	} else {
		// TODO use a sync.Pool?
		buf := make([]byte, fe.size)
		reflect.Copy(reflect.ValueOf(buf), v)
		e.Write(buf)
	}
}

type mapEncoder struct {
	encodeElem encoderFunc
}

func (me mapEncoder) encode(e *encodeState, v reflect.Value) {
	e.writeLong(int64(v.Len()))
	if sortMapKeys {
		keys := make([]string, 0, v.Len())
		for iter := v.MapRange(); iter.Next(); {
			keys = append(keys, iter.Key().String())
		}
		sort.Strings(keys)
		for _, k := range keys {
			kv := reflect.ValueOf(k)
			stringEncoder(e, kv)
			me.encodeElem(e, v.MapIndex(kv))
		}
	} else {
		for iter := v.MapRange(); iter.Next(); {
			stringEncoder(e, iter.Key())
			me.encodeElem(e, iter.Value())
		}
	}
	e.writeLong(0)
}

type arrayEncoder struct {
	encodeElem encoderFunc
}

func (ae arrayEncoder) encode(e *encodeState, v reflect.Value) {
	n := v.Len()
	e.writeLong(int64(n))
	for i := 0; i < n; i++ {
		ae.encodeElem(e, v.Index(i))
	}
	e.writeLong(0)
}

func boolEncoder(e *encodeState, v reflect.Value) {
	if v.Bool() {
		e.WriteByte(1)
	} else {
		e.WriteByte(0)
	}
}

func longEncoder(e *encodeState, v reflect.Value) {
	e.writeLong(v.Int())
}

func (e *encodeState) writeLong(x int64) {
	n := binary.PutVarint(e.scratch[:], x)
	e.Write(e.scratch[:n])
}

func floatEncoder(e *encodeState, v reflect.Value) {
	binary.LittleEndian.PutUint32(e.scratch[:], math.Float32bits(float32(v.Float())))
	e.Write(e.scratch[:4])
}

func doubleEncoder(e *encodeState, v reflect.Value) {
	binary.LittleEndian.PutUint64(e.scratch[:], math.Float64bits(v.Float()))
	e.Write(e.scratch[:8])
}

func bytesEncoder(e *encodeState, v reflect.Value) {
	data := v.Bytes()
	e.writeLong(int64(len(data)))
	e.Write(data)
}

func stringEncoder(e *encodeState, v reflect.Value) {
	s := v.String()
	e.writeLong(int64(len(s)))
	e.WriteString(s)
}

type structEncoder struct {
	fields []encoderFunc
}

func (se structEncoder) encode(e *encodeState, v reflect.Value) {
	for i, enc := range se.fields {
		enc(e, v.Field(i))
	}
}

type unionEncoderChoice struct {
	typ reflect.Type
	enc encoderFunc
}

type unionEncoder struct {
	// nullIndex holds the union index of the null alternative,
	// or -1 if there is none.
	nullIndex int
	// use a slice because unions are usually small and
	// a linear traversal is faster then.
	choices []unionEncoderChoice
}

func (ue unionEncoder) encode(e *encodeState, v reflect.Value) {
	if v.IsNil() {
		if ue.nullIndex != -1 {
			e.writeLong(int64(ue.nullIndex))
			return
		}
		e.error(fmt.Errorf("nil value not allowed"))
	}
	v = v.Elem()
	vt := v.Type()
	for i, choice := range ue.choices {
		if choice.typ == vt {
			e.writeLong(int64(i))
			choice.enc(e, v)
			return
		}
	}
	e.error(fmt.Errorf("unknown type for union %s", vt))
}

type ptrUnionEncoder struct {
	indexes    [2]byte
	encodeElem encoderFunc
}

func (pe ptrUnionEncoder) encode(e *encodeState, v reflect.Value) {
	if v.IsNil() {
		e.writeLong(int64(pe.indexes[0]))
		return
	}
	e.writeLong(int64(pe.indexes[1]))
	pe.encodeElem(e, v.Elem())
}