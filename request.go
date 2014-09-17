package goutils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
)

type RequestStruct struct {
	*http.Request
	Form *GetStruct
}

func Request(r *http.Request) *RequestStruct {
	req := RequestStruct{r, Getter(r.Form)}
	return &req
}

func (r *RequestStruct) FormatBody(v interface{}) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	switch reflect.ValueOf(v).Elem().Kind() {
	case reflect.Struct:
		return ToStruct(buf.Bytes(), v)
	default:
		if err := json.Unmarshal(buf.Bytes(), v); err != nil {
			return err
		}
	}

	return nil
}
