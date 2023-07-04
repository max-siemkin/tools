package reflect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"text/template"
)

type Map map[string]any

func (m Map) Bool(key string) bool {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case bool:
			return v
		case int, int8, int16, int32, int64:
			return m.Int(key) == 1
		case uint, uint8, uint16, uint32, uint64:
			return m.Uint(key) == 1
		case float32, float64:
			return m.Float(key) == 1
		case string:
			if b, err := strconv.ParseBool(v); err == nil {
				return b
			}
		}
	}
	return false
}
func (m Map) Int(key string) int64 {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case bool:
			if v {
				return int64(1)
			}
			return int64(0)
		case int:
			return int64(v)
		case int8:
			return int64(v)
		case int16:
			return int64(v)
		case int32:
			return int64(v)
		case int64:
			return v
		case uint:
			return int64(v)
		case uint8:
			return int64(v)
		case uint16:
			return int64(v)
		case uint32:
			return int64(v)
		case uint64:
			return int64(v)
		case float32:
			return int64(v)
		case float64:
			return int64(v)
		}
	}
	return 0
}
func (m Map) Uint(key string) uint64 {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case bool:
			if v {
				return uint64(1)
			}
			return uint64(0)
		case int:
			return uint64(v)
		case int8:
			return uint64(v)
		case int16:
			return uint64(v)
		case int32:
			return uint64(v)
		case int64:
			return uint64(v)
		case uint:
			return uint64(v)
		case uint8:
			return uint64(v)
		case uint16:
			return uint64(v)
		case uint32:
			return uint64(v)
		case uint64:
			return v
		case float32:
			return uint64(v)
		case float64:
			return uint64(v)
		}
	}
	return 0
}
func (m Map) Float(key string) float64 {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case bool:
			if v {
				return float64(1)
			}
			return float64(0)
		case int:
			return float64(v)
		case int8:
			return float64(v)
		case int16:
			return float64(v)
		case int32:
			return float64(v)
		case int64:
			return float64(v)
		case uint:
			return float64(v)
		case uint8:
			return float64(v)
		case uint16:
			return float64(v)
		case uint32:
			return float64(v)
		case uint64:
			return float64(v)
		case float32:
			return float64(v)
		case float64:
			return v
		}
	}
	return 0
}
func (m Map) Complex(key string) complex128 {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case complex64:
			return complex128(v)
		case complex128:
			return v
		}
	}
	return 0
}
func (m Map) String(key string, args ...any) string {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case bool:
			return strconv.FormatBool(v)
		case int, int8, int16, int32, int64:
			return strconv.FormatInt(m.Int(key), 10)
		case uint, uint8, uint16, uint32, uint64:
			return strconv.FormatUint(m.Uint(key), 10)
		case float32, float64:
			return strconv.FormatFloat(m.Float(key), 'E', -1, 64)
		case string:
			if c := len(args); c == 1 {
				if t, err := template.New("").Parse(v); err == nil {
					buf := bytes.NewBuffer(nil)
					if err := t.Execute(buf, args[0]); err == nil {
						return buf.String()
					}
				}
			} else if c > 1 {
				return fmt.Sprintf(v, args...)
			}
			return v
		}
	}
	return ""
}
func (m Map) Any(key string) any {
	if val, ok := m[key]; ok {
		return val
	}
	return nil
}
func (m Map) Map(key string) Map {
	if val, ok := m[key]; ok {
		if v, ok := val.(map[string]any); ok {
			return v
		}
	}
	return nil
}
func (m Map) MapArray(key string) []Map {
	maps := []Map{}
	for _, val := range m.Array(key) {
		if v, ok := val.(map[string]any); ok {
			maps = append(maps, v)
		}
	}
	return maps
}
func (m Map) Array(key string) []any {
	if val, ok := m[key]; ok {
		if v, ok := val.([]any); ok {
			return v
		}
	}
	return nil
}
func (m Map) Exist(key string) bool {
	_, ok := m[key]
	return ok
}
func (m Map) JSON(key string) string {
	if val, ok := m[key]; ok {
		data, _ := json.Marshal(&val)
		return string(data)
	}
	return ""
}
func (m Map) JSONb(key string) []byte {
	if val, ok := m[key]; ok {
		data, _ := json.Marshal(&val)
		return data
	}
	return nil
}
func (m Map) ToJSON() string {
	data, _ := json.Marshal(&m)
	return string(data)
}
func (m Map) ToJSONb() []byte {
	data, _ := json.Marshal(&m)
	return data
}
