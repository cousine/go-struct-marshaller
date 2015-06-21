package structmarshaller

import (
	"reflect"
	"strings"
)

type StructMarshaller struct {
	object interface{}
}

func New(v interface{}) *StructMarshaller {
	return &StructMarshaller{
		object: v,
	}
}

func (st *StructMarshaller) Encode() (result map[string]interface{}) {
	objReflect := reflect.TypeOf(st.object).Elem()
	objValReflect := reflect.ValueOf(st.object).Elem()
	numFields := objReflect.NumField()

	result = make(map[string]interface{})

	for i := 0; i < numFields; i++ {
		field := objReflect.Field(i)

		tag := field.Tag.Get("stmarsh")
		properties := strings.Split(tag, ",")

		key := properties[0]

		value := objValReflect.Field(i).Interface()
		zeroValue := reflect.Zero(reflect.TypeOf(value)).Interface()

		if len(properties) > 1 {
			if properties[1] == "omitempty" && value == zeroValue {
				continue
			}
		}

		result[key] = value
	}

	return
}
