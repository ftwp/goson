package goson

import (
	"encoding/json"
	"reflect"
	"unicode"
)

func UnmarshalJSON(data []byte, v interface{}) error {
	_, err := unmarshalJSON(data, reflect.ValueOf(v))
	return err
}

func unmarshalJSON(data []byte, value reflect.Value) (reflect.Value, error) {
	g, err := unmarshalGOSON(data)
	if err != nil {
		return value, err
	}
	//value := reflect.ValueOf(v)
	sourceType := value.Type()
	var e reflect.Value
	if sourceType.Kind() != reflect.Pointer {
		e = value
	} else {
		if value.IsNil() {
			value = reflect.New(sourceType.Elem())
		}
		e = value.Elem()

	}

	p := e.Type()
	fields := e.NumField()
	for i := 0; i < fields; i++ {
		jsonTag := p.Field(i).Tag.Get("json")
		if jsonTag == "-" {
			continue
		}
		name := p.Field(i).Name
		if unicode.IsUpper([]rune(name)[0]) {
			val, ok := g.v[jsonTag]
			if !ok {
				continue
			}
			elem := e.Field(i)
			v, err := unmarshalCase(elem, val)
			if err != nil {
				return v, err
			}
			elem.Set(v)
			continue
		}
		if jsonTag == "" {
			jsonTag = name
		}
		funcName := p.Field(i).Tag.Get("json-setter")
		if funcName == "" {
			funcName = "Set" + functionName(name)
		}
		val, ok := g.v[jsonTag]
		if !ok {
			continue
		}
		v, err := unmarshalCase(e.Field(i), val)
		if err != nil {
			return v, err
		}
		method := value.MethodByName(funcName)
		if !method.IsValid() {
			continue
		}
		method.Call([]reflect.Value{v})
	}
	return value, nil
}

func unmarshalCase(value reflect.Value, val json.RawMessage) (reflect.Value, error) {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var v int64
		if err := json.Unmarshal(val, &v); err != nil {
			return reflect.ValueOf(v), err
		}
		return reflect.ValueOf(v).Convert(value.Type()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var v uint64
		if err := json.Unmarshal(val, &v); err != nil {
			return reflect.ValueOf(v), err
		}
		return reflect.ValueOf(v).Convert(value.Type()), nil
	case reflect.Float32, reflect.Float64:
		var v float64
		if err := json.Unmarshal(val, &v); err != nil {
			return reflect.ValueOf(v), err
		}
		return reflect.ValueOf(v).Convert(value.Type()), nil
	case reflect.Complex64, reflect.Complex128:
		var v complex128
		if err := json.Unmarshal(val, &v); err != nil {
			return reflect.ValueOf(v), err
		}
		return reflect.ValueOf(v).Convert(value.Type()), nil
	case reflect.Array:
		var v any
		if err := json.Unmarshal(val, &v); err != nil {
			return reflect.ValueOf(&v), err
		}
		return reflect.ValueOf(v), nil
	case reflect.String:
		var v string
		if err := json.Unmarshal(val, &v); err != nil {
			return reflect.ValueOf(v), err
		}
		return reflect.ValueOf(v).Convert(value.Type()), nil
	case reflect.Pointer:
		v := reflect.New(value.Type().Elem())
		return unmarshalJSON(val, v)
	default:
		v := reflect.New(value.Type())
		return unmarshalJSON(val, v.Elem())
	}
}
