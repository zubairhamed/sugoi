package sugoi

import (
	"net/http"
	"strconv"
)

func NewRequestFromHttp(attrs map[string]string, req *http.Request) *Request {
	return &Request{
		attrs:       attrs,
		httpRequest: req,
	}
}

type Request struct {
	attrs       map[string]string
	httpRequest *http.Request
	// Query
	// Cookies
	// Sessions
	// get attribute
	// set attribute
}

func (c *Request) GetHttpRequest() *http.Request {
	return c.httpRequest
}

func (c *Request) GetAttributes() map[string]string {
	return c.attrs
}

func (c *Request) GetAttribute(o string) string {
	return c.attrs[o]
}

func (c *Request) GetAttributeAsInt(o string) int {
	attr := c.GetAttribute(o)
	i, _ := strconv.Atoi(attr)

	return i
}

func NewWrappedHandler() *WrappedHandler {
	return &WrappedHandler{
		routes:          []*Route{},
		preFilters:      []PreFilter{},
		defaultHandlers: make(map[int]RouteHandler),
		staticUrl:       "/static",
		staticDir:       "static",
	}
}
