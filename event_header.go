// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCES:
 *     DefaultEventCreated.avsc
 *     EventHeader.avsc
 *     NewEvent2.avsc
 */
package my_sandbox_sandbox_events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

var _ = fmt.Printf

// The below fields include header information and should be included on every event in the DESP. Inspired by: https://github.com/cloudevents/spec/blob/v0.2/spec.md
type EventHeader struct {
	// A unique identifier of the event - for example, a randomly generated GUID
	Id string `json:"id"`
	// Time the event occurred in milliseconds since epoch, UTC timezone.
	Time int64 `json:"time"`
	// Type of occurrence which has happened. Reference the domain.event registered in schema-registry.
	Type string `json:"type"`
	// Service that produced the event. Future: reference to producer registry.
	Source string `json:"source"`
}

const EventHeaderAvroCRC64Fingerprint = "\xb6\x98w\xe1\xb8\u007f\x85\xd8"

func NewEventHeader() EventHeader {
	r := EventHeader{}
	return r
}

func DeserializeEventHeader(r io.Reader) (EventHeader, error) {
	t := NewEventHeader()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventHeaderFromSchema(r io.Reader, schema string) (EventHeader, error) {
	t := NewEventHeader()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventHeader(r EventHeader, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Time, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Source, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventHeader) Serialize(w io.Writer) error {
	return writeEventHeader(r, w)
}

func (r EventHeader) Schema() string {
	return "{\"doc\":\"The below fields include header information and should be included on every event in the DESP. Inspired by: https://github.com/cloudevents/spec/blob/v0.2/spec.md\",\"fields\":[{\"doc\":\"A unique identifier of the event - for example, a randomly generated GUID\",\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Time the event occurred in milliseconds since epoch, UTC timezone.\",\"name\":\"time\",\"type\":\"long\"},{\"doc\":\"Type of occurrence which has happened. Reference the domain.event registered in schema-registry.\",\"name\":\"type\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Service that produced the event. Future: reference to producer registry.\",\"name\":\"source\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}}],\"name\":\"com.kroger.desp.commons.desp.sandbox.EventHeader\",\"type\":\"record\"}"
}

func (r EventHeader) SchemaName() string {
	return "com.kroger.desp.commons.desp.sandbox.EventHeader"
}

func (_ EventHeader) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventHeader) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventHeader) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventHeader) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventHeader) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventHeader) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventHeader) SetString(v string)   { panic("Unsupported operation") }
func (_ EventHeader) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventHeader) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Id}
	case 1:
		return &types.Long{Target: &r.Time}
	case 2:
		return &types.String{Target: &r.Type}
	case 3:
		return &types.String{Target: &r.Source}
	}
	panic("Unknown field index")
}

func (r *EventHeader) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventHeader) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventHeader) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventHeader) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventHeader) Finalize()                        {}

func (_ EventHeader) AvroCRC64Fingerprint() []byte {
	return []byte(EventHeaderAvroCRC64Fingerprint)
}

func (r EventHeader) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["time"], err = json.Marshal(r.Time)
	if err != nil {
		return nil, err
	}
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	output["source"], err = json.Marshal(r.Source)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventHeader) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["time"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Time); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for time")
	}
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	val = func() json.RawMessage {
		if v, ok := fields["source"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Source); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for source")
	}
	return nil
}
