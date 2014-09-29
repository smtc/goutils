package goutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/guotie/config"
)

var (
	templatePath = config.GetStringDefault("templatePath", "./assets/templates")
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, ".html") {
		Render(w).RenderHtml(r.URL.Path)
		return
	}

	http.Error(w, fmt.Sprintf("%s page not found.", r.URL.Path), 404)
}

// ==============================================================

type RenderStruct struct {
	http.ResponseWriter
}

func Render(w http.ResponseWriter) *RenderStruct {
	return &RenderStruct{w}
}

func (w *RenderStruct) RenderHtml(path string) {
	buf, err := ioutil.ReadFile(templatePath + path)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(buf)
}

func (w *RenderStruct) RenderPage(v interface{}, total int, r *http.Request) {
	type Page struct {
		Status  int         `json:"status"`
		Data    interface{} `json:"data"`
		Total   int         `json:"total"`
		hasNext bool        `json:"hasnext"`
		size    int         `json:"size"`
		page    int         `json:"page"`
	}
	page := Page{
		Status: 1,
		Total:  total,
		Data:   v,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	buf, _ := json.Marshal(page)
	w.Write(buf)
}

func (w *RenderStruct) RenderJson(v interface{}, status int) {
	type Result struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
	}
	result := Result{
		Status: status,
		Data:   v,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	buf, _ := json.Marshal(result)
	w.Write(buf)
}

func (w *RenderStruct) RenderJsonNoWrap(v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	buf, _ := json.Marshal(v)
	w.Write(buf)
}

func (w *RenderStruct) RenderError(err string) {
	type Result struct {
		Status int    `json:"status"`
		Errors string `json:"errors"`
	}
	result := Result{
		Status: 0,
		Errors: err,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	buf, _ := json.Marshal(result)
	w.Write(buf)
}
