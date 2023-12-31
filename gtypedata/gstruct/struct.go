package gstruct

import (
	"reflect"
)

// ExtractStructTagsAndFields takes a struct and returns a map of its field names
// and their respective tags, excluding those tagged with "ignore". It handles nested
// struct fields designated with the tag "-".
// The map result is a key value pair of field name and tag. ex : {"ID":"id|int64","Name":"name|string"}
// Nested is not supported.
func ExtractStructTagsAndFields(src any, prefix string, tag string) (s map[string]string) {
	var val reflect.Value

	if reflect.ValueOf(src).Kind() == reflect.Ptr {
		val = reflect.ValueOf(src).Elem()
	} else {
		val = reflect.ValueOf(src)
	}

	typ := val.Type()
	s = make(map[string]string)
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)

		// Skip fields with tag "ignore"
		if field.Tag.Get(tag) == "ignore" {
			continue
		}

		fieldTag := field.Tag.Get(tag)
		switch val.Field(i).Kind() {
		case reflect.Struct:
			if field.Type.String() == "gdb.SqlNullString" ||
				field.Type.String() == "gdb.SqlNullBool" ||
				field.Type.String() == "gdb.SqlNullFloat64" ||
				field.Type.String() == "gdb.SqlNullInt64" ||
				field.Type.String() == "gdb.SqlNullInt32" ||
				field.Type.String() == "gdb.SqlNullByte" ||
				field.Type.String() == "gdb.SqlNullTime" ||
				field.Type.String() == "time.Time" ||
				field.Type.String() == "gdb.SqlNullInt16" {
				if fieldTag != "" {
					s[field.Name] = fieldTag + "|" + field.Type.String()
				}
				continue
			}

			if field.Tag.Get(tag) != "-" && field.Tag.Get(tag) != "ignore" && field.Tag.Get(tag) != "" {
				panic("nested struct is not supported")
			}
			if fieldTag == "-" {
				res := ExtractStructTagsAndFields(val.Field(i).Interface(), prefix, tag)
				for k, v := range res {
					s[k] = v
				}
			} else {
				res := ExtractStructTagsAndFields(val.Field(i).Interface(), fieldTag+".", tag)
				for k, v := range res {
					s[k] = v
				}
			}
		default:
			if fieldTag == "-" {
				continue
			}
			if fieldTag != "" {
				s[field.Name] = fieldTag + "|" + field.Type.String()
			}
		}
	}

	return s
}
