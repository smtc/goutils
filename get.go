/*
* get value from struct or map by string key
* if key does not exist, return given default value
 */

package goutils

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type GetStruct struct {
	Value interface{}
}

func Getter(v interface{}) *GetStruct {
	return &GetStruct{Value: v}
}

func (g *GetStruct) GetValue(key string) reflect.Value {
	fv := reflect.ValueOf(g.Value)
	switch fv.Kind() {

	case reflect.Struct:
		return fv.FieldByName(key)

	case reflect.Map:
		return fv.MapIndex(reflect.ValueOf(key))

	}

	return fv
}

func (g *GetStruct) GetInterface(key string, def interface{}) interface{} {
	val := g.GetValue(key)
	if !val.IsValid() {
		return def
	}

	return val.Interface()
}

func ToInt64(v interface{}, def int64) int64 {
	if i, ok := v.(int64); ok {
		return i
	}
	if i, ok := v.(int); ok {
		return int64(i)
	}
	if i, ok := v.(float32); ok {
		return int64(i)
	}
	if i, ok := v.(float64); ok {
		return int64(i)
	}
	if b, ok := v.([]byte); ok {
		return int64(binary.BigEndian.Uint64(b))
	}
	if ss, ok := v.([]string); ok {
		v = ss[0]
	}
	if s, ok := v.(string); ok {
		if i, err := strconv.ParseInt(s, 0, 64); err == nil {
			return i
		}
	}
	return def
}

func (g *GetStruct) GetInt64(key string, def int64) int64 {
	v := g.GetInterface(key, def)
	return ToInt64(v, def)
}

func ToInt(v interface{}, def int) int {
	return int(ToInt64(v, int64(def)))
}

func (g *GetStruct) GetInt(key string, def int) int {
	i := g.GetInt64(key, int64(def))
	return int(i)
}

func ToFloat64(v interface{}, def float64) float64 {
	if f, ok := v.(float64); ok {
		return f
	}
	if ss, ok := v.([]string); ok {
		v = ss[0]
	}
	if s, ok := v.(string); ok {
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return f
		}
	}
	return def
}

func (g *GetStruct) GetFloat64(key string, def float64) float64 {
	v := g.GetInterface(key, def)
	return ToFloat64(v, def)
}

func ToFloat32(v interface{}, def float32) float32 {
	return float32(ToFloat64(v, float64(def)))
}

func (g *GetStruct) GetFloat32(key string, def float32) float32 {
	f := g.GetFloat64(key, float64(def))
	return float32(f)
}

func ToString(v interface{}, def string) string {
	if s, ok := v.(string); ok {
		return s
	}
	if ss, ok := v.([]string); ok {
		return strings.Join(ss, ",")
	}

	fv := reflect.ValueOf(v)
	if fv.Kind() == reflect.Slice {
		var ss []string
		for i := 0; i < fv.Len(); i++ {
			ss = append(ss, fmt.Sprintf("%v", fv.Index(i).Interface()))
		}
		return strings.Join(ss, ",")
	}

	if t, ok := v.(time.Time); ok {
		return t.Format(TIMEFORMAT)
	}

	return fmt.Sprintf("%v", v)
}

func (g *GetStruct) GetString(key string, def string) string {
	v := g.GetInterface(key, def)
	return ToString(v, def)
}

func ToTime(v interface{}, def time.Time, ft string) time.Time {
	if t, ok := v.(time.Time); ok {
		return t
	}
	if ss, ok := v.([]string); ok {
		v = ss[0]
	}
	if s, ok := v.(string); ok {
		if t, err := time.Parse(ft, s); err == nil {
			return t
		}
	}

	return def
}

func (g *GetStruct) GetTime(key string, def time.Time, ft string) time.Time {
	v := g.GetInterface(key, def)
	return ToTime(v, def, ft)
}

// strconv.ParseBool can parse 1 & 0
// when v is string "", return false
func ToBool(v interface{}, def bool) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	if i, ok := v.(int); ok {
		return i > 0
	}
	if i, ok := v.(float64); ok {
		return i > 0
	}
	if i, ok := v.(float32); ok {
		return i > 0
	}
	if ss, ok := v.([]string); ok {
		v = ss[0]
	}
	if s, ok := v.(string); ok {
		if s == "on" {
			return true
		}
		if s == "off" || s == "" {
			return false
		}
		if b, err := strconv.ParseBool(s); err == nil {
			return b
		}
	}

	return def

}

func (g *GetStruct) GetBool(key string, def bool) bool {
	v := g.GetInterface(key, def)
	return ToBool(v, def)
}

func ToBytes(v interface{}, def []byte) []byte {
	if b, ok := v.([]byte); ok {
		return b
	}
	if ss, ok := v.([]string); ok {
		v = ss[0]
	}
	if s, ok := v.(string); ok {
		return []byte(s)
	}

	return def
}

func (g *GetStruct) GetBytes(key string, def []byte) []byte {
	v := g.GetInterface(key, def)

	return ToBytes(v, def)
}
