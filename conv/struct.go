package conv

import (
	"reflect"
)

// StructToMap :
// 结构体对象转化为集合
func StructToMap(obj interface{}) map[string]interface{} {
	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()
	var data = make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {
		data[typ.Field(i).Name] = val.Field(i).Interface()
	}
	return data
}
