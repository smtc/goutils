package goutils

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

type httpHandler struct {
	Env      *GetStruct
	Param    *GetStruct
	Query    *GetStruct
	Response *RenderStruct
	Request  *RequestStruct
}

func HttpHandler(c web.C, w http.ResponseWriter, r *http.Request) *httpHandler {
	return &httpHandler{
		Env:      Getter(c.Env),
		Param:    Getter(c.URLParams),
		Query:    Getter(r.URL.Query()),
		Response: Render(w),
		Request:  Request(r),
	}
}

func (h *httpHandler) RenderHtml(path string) {
	h.Response.RenderHtml(path)
}

func (h *httpHandler) RenderPage(v interface{}) {
	h.Response.RenderPage(v)
}

func (h *httpHandler) RenderJson(v interface{}, status int) {
	h.Response.RenderJson(v, status)
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
