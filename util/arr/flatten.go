package arr

import "reflect"

func Flatten(args ...interface{}) (res []interface{}) {
	return flatten(reflect.ValueOf(args))
}

func flatten(args reflect.Value) (res []interface{}) {
	for idx, l := 0, args.Len(); idx < l; idx++ {
		val := args.Index(idx)
		kind := val.Kind()
		if kind == reflect.Interface {
			val = val.Elem()
			kind = val.Kind()
		}
		if kind == reflect.Array || kind == reflect.Slice {
			res = append(res, flatten(val)...)
		} else {
			res = append(res, val.Interface())
		}
	}

	return
}
