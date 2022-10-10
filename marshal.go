package goson

import (
	"encoding/json"
	"errors"
	"reflect"
	"unicode"
)

func MarshalJSON(v interface{}) ([]byte, error) {
	g := emptyGOSON()
	value := reflect.ValueOf(v)
	var e reflect.Value
	if value.Type().Kind() == reflect.Struct {
		e = value
	} else {
		if value.IsNil() {
			return nil, errors.New("empty interface")
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
			//fmt.Println("title: ", name)
			elemType := p.Field(i).Type
			elem := e.Field(i)
			g.v[jsonTag], _ = marshalCase(elemType, elem)
			continue
		}
		if jsonTag == "" {
			jsonTag = name
		}
		funcName := p.Field(i).Tag.Get("json-getter")
		if funcName == "" {
			funcName = functionName(name)
		}
		//fmt.Println("funcName:", funcName)
		method := value.MethodByName(funcName)
		if !method.IsValid() {
			continue
		}
		elems := method.Call(nil)
		if len(elems) == 0 {
			continue
		}
		elemType := elems[0].Type()
		g.v[jsonTag], _ = marshalCase(elemType, elems[0])
	}
	//fmt.Println("fields: ", fields)
	//fmt.Println("element.string: ", elem.String())
	//fmt.Println("element.type: ", elem.Type())
	//fmt.Println("element.type: ")
	return g.MarshalJSON()
}

func marshalCase(p reflect.Type, value reflect.Value) ([]byte, error) {
	//fmt.Println("case type: ", p, "value: ", value)
	switch p.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return json.Marshal(value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return json.Marshal(value.Uint())
	case reflect.Float32, reflect.Float64:
		return json.Marshal(value.Float())
	case reflect.Complex64, reflect.Complex128:
		return json.Marshal(value.Complex())
	case reflect.Array:
		return json.Marshal(value.Interface())
	case reflect.String:
		return json.Marshal(value.String())
	default:

	}
	return MarshalJSON(value.Interface())
}

func functionName(name string) string {
	nameRunes := []rune(name)
	var ret []rune
	if nameRunes[0] >= 0x61 && nameRunes[0] <= 0x7A {
		ret = append(ret, unicode.ToUpper(nameRunes[0]))
		nameRunes = nameRunes[1:]
	}
	nextUpper := false
	for i := range nameRunes {
		if nameRunes[i] == rune('_') {
			nextUpper = true
			continue
		}
		if nextUpper {
			nextUpper = false
			ret = append(ret, unicode.ToUpper(nameRunes[i]))
			continue
		}
		ret = append(ret, nameRunes[i])
	}
	return string(ret)
}
