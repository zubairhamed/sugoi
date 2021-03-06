package sugoi

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func CreateCompilableRoutePath(route string) (*regexp.Regexp, bool) {
	var re *regexp.Regexp
	var isStatic bool

	regexpString := route

	isStaticRegexp := regexp.MustCompile(`[\(\)\?\<\>:]`)
	if !isStaticRegexp.MatchString(route) {
		isStatic = true
	}

	// Dots
	re = regexp.MustCompile(`([^\\])\.`)
	regexpString = re.ReplaceAllStringFunc(regexpString, func(m string) string {
		return fmt.Sprintf(`%s\.`, string(m[0]))
	})

	// Wildcard names
	re = regexp.MustCompile(`:[^/#?()\.\\]+\*`)
	regexpString = re.ReplaceAllStringFunc(regexpString, func(m string) string {
		return fmt.Sprintf("(?P<%s>.+)", m[1:len(m)-1])
	})

	re = regexp.MustCompile(`:[^/#?()\.\\]+`)
	regexpString = re.ReplaceAllStringFunc(regexpString, func(m string) string {
		return fmt.Sprintf(`(?P<%s>[^/#?]+)`, m[1:len(m)])
	})

	s := fmt.Sprintf(`\A%s\z`, regexpString)

	return regexp.MustCompile(s), isStatic
}

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
	path    string
	method  string
	handler RouteHandler
	regEx   *regexp.Regexp
}

func MatchingRoute(path string, method string, routes []*Route) (RouteHandler, map[string]string, error) {
	for _, route := range routes {
		if method == route.method {
			match, attrs := MatchesRoutePath(path, route.regEx)
			if match {
				return route.handler, attrs, nil
			}
		}
	}
	return nil, nil, ERR_NO_MATCHING_ROUTE
}

type WrappedHandler struct {
	routes          []*Route
	preFilters      []PreFilter
	errorHandlers   []ErrorHandler
	defaultHandlers map[int]RouteHandler

	staticUrl string
	staticDir string
}

func invokePreFilters(filters []PreFilter, req *Request) *Request {
	if len(filters) > 0 {
		chain := NewPreFilterChain(filters)
		if chain != nil {
			chain.filter.(PreFilter)(req, chain)

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
		content:  content,
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
	method := strings.ToLower(r.Method)

	if wh.staticUrl != "" {
		if strings.HasPrefix(urlPath, wh.staticUrl) {
			http.ServeFile(w, r, urlPath[1:])
			return
		}
	}

	fn, attrs, err := MatchingRoute(urlPath, method, wh.routes)
	req := NewRequestFromHttp(attrs, r)
	req = invokePreFilters(wh.preFilters, req)

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
