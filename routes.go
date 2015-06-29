package sugoi
import (
	"regexp"
	"fmt"
	"net/http"
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
	routes 	[]*Route
}

func (wh *WrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Match Before

	fn, attrs, err := MatchingRoute(r.URL.Path, r.Method, wh.routes)

	if err != nil {
		if err == ERR_NO_MATCHING_ROUTE {
			http.NotFound(w, r)
			return
		}
	} else {
		req := NewRequestFromHttp(attrs)
		resp := fn(req)

		SendHttpResponse(resp, w)
		return
	}

	// TODO: Match After
}
