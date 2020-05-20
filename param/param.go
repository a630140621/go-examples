package param

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

// 类似 json.Unmarshal
func Unmarshal(data []byte, ptr interface{}) error {
	// return json.Unmarshal(data, ptr)
	jsonData, err := getJSONData(data)
	if err != nil {
		return errors.Wrap(err, "json.Unmarshal err")
	}

	fmt.Printf("json data %+v\n", jsonData)

	return unmarshal(jsonData, ptr)
}

func unmarshal(jsonData *map[string]interface{}, ptr interface{}) error {
	rv := reflect.ValueOf(ptr).Elem()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fmt.Printf("field kind = %q\n", field.Kind().String())
		fieldInfo := rv.Type().Field(i)
		fmt.Printf("struct field name = %q, type = %q, tag = %q, anonymous = %t\n", fieldInfo.Name, fieldInfo.Type, fieldInfo.Tag, fieldInfo.Anonymous)

		switch field.Kind() {
		case reflect.Struct:
			if fieldInfo.Anonymous {
				unmarshal(jsonData, field.Addr().Interface())
			} else {
				d := (*jsonData)[fieldInfo.Name].(map[string]interface{})
				unmarshal(&d, field.Addr().Interface())
			}
		case reflect.String, reflect.Bool:
			field.Set(reflect.ValueOf((*jsonData)[fieldInfo.Name]))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(int64(reflect.ValueOf((*jsonData)[fieldInfo.Name]).Float()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(uint64(reflect.ValueOf((*jsonData)[fieldInfo.Name]).Uint()))
		case reflect.Float32, reflect.Float64:
			f := (*jsonData)[fieldInfo.Name].(float64)
			field.SetFloat(f)
		default:
			return fmt.Errorf("do not support this type %q", field.Kind())
		}
	}

	return nil
}

func getJSONData(data []byte) (*map[string]interface{}, error) {
	ret := map[string]interface{}{}
	// json.Unmarshal 解码后对应类型
	// Bool                   对应JSON布尔类型
	// float64                对应JSON数字类型
	// string                 对应JSON字符串类型
	// []interface{}          对应JSON数组
	// map[string]interface{} 对应JSON对象
	// nil                    对应JSON的null
	err := json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
