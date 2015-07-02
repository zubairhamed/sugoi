package sugoi
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	re, static := CreateCompilableRoutePath("/")
	assert.True(t, static)
	assert.True(t, re.MatchString("/"))

	re, static = CreateCompilableRoutePath("/test")
	assert.True(t, static)
	assert.False(t, re.MatchString("/"))
	assert.True(t,  re.MatchString("/test"))

	re, static = CreateCompilableRoutePath("/test/:var")
	assert.False(t, static)
	assert.True(t, re.MatchString("/test/abc"))
	assert.False(t, re.MatchString("/test/abc/def"))

	fn := func(*Request) Content { return nil }
	// CreateNewRoute
	route := CreateNewRoute("/", "GET", fn)

	assert.NotNil(t, route)

	assert.Equal(t, "GET", route.method)
	assert.Equal(t, "/", route.path)
	assert.True(t, route.regEx.MatchString("/"))
	assert.False(t, route.regEx.MatchString("/test"))
}

/*
func CreateNewRoute(path string, method string, fn RouteHandler) *Route {
	re, _ := CreateCompilableRoutePath(path)

	return &Route{
		path:    path,
		method:  method,
		handler: fn,
		regEx:   re,
	}
}

func MatchesRoutePath(path string, re *regexp.Regexp) (bool, map[string]string) {
	matches := re.FindAllStringSubmatch(path, -1)
	attrs := make(map[string]string)
	if len(matches) > 0 {
		subExp := re.SubexpNames()
		for idx, exp := range subExp {
			attrs[exp] = matches[0][idx]
		}
		return true, attrs
	}
	return false, attrs
}

type Route struct {
	path 	string
	method 	string
	handler RouteHandler
	regEx 	*regexp.Regexp
}

func MatchingRoute(path string, method string, routes []*Route) (RouteHandler, map[string]string, error) {
	for _, route := range routes {
		match, attrs :=  MatchesRoutePath(path, route.regEx)
		if match {
			return route.handler, attrs, nil
		}
	}
	return nil, nil, ERR_NO_MATCHING_ROUTE
}

type WrappedHandler struct {
	routes 			[]*Route
	beforeFilters	[]BeforeFilter
	errorHandlers	[]ErrorHandler
	defaultHandlers map[int]RouteHandler

	staticUrl 		string
	staticDir 		string
}

func invokeBeforeFilters(filters []BeforeFilter, req *Request) *Request {
	if len(filters) > 0 {
		chain := NewBeforeFilterChain(filters)
		if chain != nil {
			chain.filter.(BeforeFilter)(req, chain)

			req = chain.GetFilteredRequest()
		}
	}
	return req
}

func (wh *WrappedHandler) CallDefaultHandler(code int, req *Request) Content {
	h := wh.defaultHandlers[code]
	if h != nil {
		return h(req)
	} else {
		switch code {
			case 404:
			return NotFound()

			case 500:
			return InternalServerError()
		}
		return InternalServerError()
	}
}

func NewResponse(content interface{}, httpCode int) *Response {
	return &Response{
		httpCode: httpCode,
		content: content,
	}
}

func SendResponse(content interface{}, w http.ResponseWriter) {
	var response *Response
	if val, ok := content.(HttpCode); ok {
		response = NewResponse(val.content, val.code)
	} else {
		response = NewResponse(content, http.StatusOK)
	}
	ResponseHandler(response, w)
}

func (wh *WrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	if strings.HasPrefix(urlPath, wh.staticUrl) {
		http.ServeFile(w, r, urlPath[1:])
		return
	}

	fn, attrs, err := MatchingRoute(urlPath, r.Method, wh.routes)
	req := NewRequestFromHttp(attrs, r)

	req = invokeBeforeFilters(wh.beforeFilters, req)

	if err != nil {
		if err == ERR_NO_MATCHING_ROUTE {
			content := wh.CallDefaultHandler(404, req)
			SendResponse(content, w)
			return
		}
	} else {
		SendResponse(fn(req), w)
	}
}

 */