package gstruct

import (
	"encoding/json"
	"reflect"
)

func MarshalAndCencoredTag(src any, tagCencored string) (string, error) {
	var val reflect.Value
	if reflect.ValueOf(src).Kind() == reflect.Ptr {
		val = reflect.ValueOf(src).Elem()
	} else {
		val = reflect.ValueOf(src)
	}
	typ := val.Type()

	result := make(map[string]any)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get(tagCencored)
		if tag == "" {
			result[field.Tag.Get("json")] = val.Field(i).Interface()
		}
	}

	marshal, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}
