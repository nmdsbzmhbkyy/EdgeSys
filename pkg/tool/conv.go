package tool

import (
	"encoding/json"
	"mod.miligc.com/edge-common/business-common/business/pkg"
	"reflect"
	"strings"
	"time"
)

// SnakeString snake string, XxYy to xx_yy , XxYY to xx_y_y
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data))
}

// CamelString camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	flag, num := true, len(s)-1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' {
			flag = true
			continue
		} else if flag {
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
			flag = false
		}
		data = append(data, d)
	}
	return string(data)
}

func FirstLowCamelString(s string) string {
	data := make([]byte, 0, len(s))
	flag, num := true, len(s)-1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d == '_' {
			flag = true
			continue
		} else if flag {
			if d >= 'a' && d <= 'z' {
				d = d - 32
			}
			flag = false
		}
		data = append(data, d)
	}
	if len(data) > 0 && data[0] >= 65 && data[0] <= 90 {
		data[0] = data[0] + 32
	}
	return string(data)
}

func StructToMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	value := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i).Interface()
		result[field.Name] = fieldValue
	}

	return result
}

func MapToStruct(m map[string]interface{}, s interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		return err
	}

	return nil
}

func InterfaceToStruct(m interface{}, s interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		return err
	}

	return nil
}

func StringToStruct(m string, s interface{}) error {
	err := json.Unmarshal([]byte(m), s)
	if err != nil {
		return err
	}
	return nil
}

func TimeToFormat(val interface{}) string {
	switch v := val.(type) {
	case int64:
		// 如果是时间戳类型，将其转换为时间对象
		t := time.Unix(v, 0)
		// 格式化时间字符串
		formattedTime := t.Format("2006-01-02 15:04:05")
		return formattedTime
	case string:
		// 如果是字符串类型，将其解析为时间对象
		t, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			pkg.Log.Error("时间格式非标准格式")
			return ""
		}
		// 格式化时间字符串
		formattedTime := t.Format("2006-01-02 15:04:05")
		return formattedTime
	default:
		return ""
	}
}
