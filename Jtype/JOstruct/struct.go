package JOstruct

import (
	"reflect"
)

// GetTagAndFieldNameFromStruct is a function to extract the tag and field name
// from a struct. If a field within the struct includes the tag "ignore", it will be
// skipped. If the tag is "-", it means that it will enter the nested struct fields
// and the prefix used here will be from the parameter, not from the tag of the nested struct.
// map result is a key value pair of field name and tag. ex : {"ID":"id|JOsql.NullString"}
func GetTagAndFieldNameFromStruct(src any, prefix string, tag string) (s map[string]string) {
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
			if field.Type.String() == "JOsql.NullString" ||
				field.Type.String() == "JOsql.NullBool" ||
				field.Type.String() == "JOsql.NullFloat64" ||
				field.Type.String() == "JOsql.NullInt64" ||
				field.Type.String() == "JOsql.NullInt32" ||
				field.Type.String() == "JOsql.NullByte" ||
				field.Type.String() == "JOsql.NullTime" ||
				field.Type.String() == "JOsql.NullInt16" {
				if fieldTag != "" {
					s[field.Name] = fieldTag + "|" + field.Type.String()
				}
				continue
			}
			if field.Tag.Get(tag) != "-" && field.Tag.Get(tag) != "ignore" {
				panic("nested struct is not supported")
			}
			if fieldTag == "-" {
				res := GetTagAndFieldNameFromStruct(val.Field(i).Interface(), prefix, tag)
				for k, v := range res {
					s[k] = v
				}
			} else {
				res := GetTagAndFieldNameFromStruct(val.Field(i).Interface(), fieldTag+".", tag)
				for k, v := range res {
					s[k] = v
				}
			}
		default:
			if fieldTag != "" {
				s[field.Name] = fieldTag + "|" + field.Type.String()
			}
		}
	}

	return s
}
