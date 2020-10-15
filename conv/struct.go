package conv

import "reflect"

// StructToMap :
// 结构体对象转化为集合
func StructToMap(obj interface{}) map[string]interface{} {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {
		data[typ.Field(i).Name] = val.Field(i).Interface()
	}
	return data
}
