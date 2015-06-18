package structmarshaller

import (
	"reflect"
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
		key := field.Tag.Get("stmarsh")
		result[key] = objValReflect.Field(i).Interface()
	}

	return
}
