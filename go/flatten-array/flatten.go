package flatten

import (
	"reflect"
)

func Flatten(array interface{}) (result []interface{}) {
	result = []interface{}{}
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice:
		for _, elem := range array.([]interface{}) {
			ve := reflect.ValueOf(elem)
			switch ve.Kind() {
			case reflect.Slice:
				result = append(result, Flatten(elem)...)
			case reflect.Int:
				result = append(result, elem)
			}
		}

	case reflect.Int:
		result = append(result, v)
	default:
		return
	}
	return
}
