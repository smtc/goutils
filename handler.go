package goutils

import (
	"net/http"
	"net/url"

	"github.com/zenazn/goji/web"
)

type httpHandler struct {
	Env      *GetStruct
	Param    *GetStruct
	Query    *GetStruct
	Form     url.Values
	Response *RenderStruct
	Request  *RequestStruct
	r        *http.Request
}

func HttpHandler(c web.C, w http.ResponseWriter, r *http.Request) *httpHandler {
	return &httpHandler{
		Env:      Getter(c.Env),
		Param:    Getter(c.URLParams),
		Query:    Getter(r.URL.Query()),
		Response: Render(w),
		Request:  Request(r),
		r:        r,
	}
}

func (h *httpHandler) RenderHtml(path string) {
	h.Response.RenderHtml(path)
}

func (h *httpHandler) RenderPage(v interface{}, total int) {
	h.Response.RenderPage(v, total, h.r)
}

func (h *httpHandler) RenderJson(v interface{}, status int, msg string) {
	h.Response.RenderJson(v, status, msg)
}

func (h *httpHandler) RenderJsonNoWrap(v interface{}) {
	h.Response.RenderJsonNoWrap(v)
}

func (h *httpHandler) RenderError(err string) {
	h.Response.RenderError(err)
}

func (h *httpHandler) FormatBody(v interface{}) error {
	return h.Request.FormatBody(v)
}

func (h *httpHandler) ParseForm() {
	h.r.ParseForm()
	h.Form = h.r.Form
}

func (h *httpHandler) GetPageSize() (int, int) {
	var (
		page int
		size int
	)
	page = h.Query.GetInt("page", 1)
	if page < 1 {
		page = 1
	}
	size = h.Query.GetInt("size", 20)
	return page, size
}
