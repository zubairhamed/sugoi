package main
import (
	"log"
	"net/http"
	"errors"
	"regexp"
	"strings"
	"strconv"
)

type Response interface{}
type RouteHandler func(Request) Response


/*
	return NotFound("Missing file")
 */
func main() {
	s := NewSugoi()

	s.GET("/", func(req Request) Response {
		return "homepage"
	})

	s.GET("/hello/{name}", func(req Request) Response {
		name := req.GetAttribute("name")
		return "hello " + name
	})


	s.DELETE("/", func(req Request) Response {
		return nil
	})

	s.PUT("/", func(req Request) Response {
		return nil
	})

	s.POST("/", func(req Request) Response {
		return nil
	})

	s.OPTIONS("/", func(req Request) Response {
		return nil
	})

	s.PATCH("/", func(req Request) Response {
		return nil
	})

	s.Serve()
}

func NewSugoi() *SugoiServer {

	return &SugoiServer{
		handler: NewWrappedHandler(),
	}
}

type SugoiServer struct {
	handler 	*WrappedHandler
}

func (s *SugoiServer) add(method string, path string, fn RouteHandler) {
	s.handler.routes = append(s.handler.routes, CreateNewRoute(path, method, fn))
}

func (s *SugoiServer) GET(path string, fn RouteHandler) {
	s.add("get", path, fn)
}

func (s *SugoiServer) DELETE(path string, fn RouteHandler) {
	s.add("delete", path, fn)
}

func (s *SugoiServer) PUT(path string, fn RouteHandler) {
	s.add("put", path, fn)
}

func (s *SugoiServer) POST(path string, fn RouteHandler) {
	s.add("post", path, fn)
}

func (s *SugoiServer) OPTIONS(path string, fn RouteHandler) {
	s.add("options", path, fn)
}

func (s *SugoiServer) PATCH(path string, fn RouteHandler) {
	s.add("patch", path, fn)
}

func (s *SugoiServer) Serve() {
	log.Println("Starting Sugoi!")

	err := http.ListenAndServe(":8085", s.handler)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateCompilableRoutePath(pre string) (*regexp.Regexp, error) {
	re, _ := regexp.Compile(`{[a-z]+}`)

	matches := re.FindAllStringSubmatch(pre, -1)

	pre = "^" + pre
	for _, b := range matches {
		origAttr := b[0]
		attr := strings.Replace(strings.Replace(origAttr, "{", "", -1), "}", "", -1)
		frag := `(?P<` + attr + `>\w+)`
		pre = strings.Replace(pre, origAttr, frag, -1)
	}
	pre += "$"
	re, err := regexp.Compile(pre)

	return re, err
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

func NewWrappedHandler() (*WrappedHandler) {
	return &WrappedHandler{
		routes : []*Route{},
	}
}

var ERR_NO_MATCHING_ROUTE = errors.New("No matching route found")
var ERR_UNSUPPORTED_CONTENT_FORMAT = errors.New("Unsupported Content-Format")
var ERR_NO_MATCHING_METHOD = errors.New("No matching method")

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
}

func SendHttpResponse(response interface{}, w http.ResponseWriter) {
	log.Println("Sending Back Response..", response)
	ResponseHandler(response, w)
}

func ResponseHandler(response interface{}, w http.ResponseWriter) {

}

func NewRequestFromHttp(attrs map[string]string) Request {
	return Request{
		attrs: attrs,
	}
}


type Request struct {
	attrs 	map[string]string
	// Query
	// Cookies
	// Sessions
		// get attribute
		// set attribute
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


// Functions
func Halt(msg string, code int) {
	// halt(int)
	// halt(msg)
	// halt(int, msg)
}

func HaltWithCode(code int) {

}

func HaltWithMessage(msg string) {

}

// Before every method
func BeforeAll(fn RouteHandler) {

}

func Before(pattern string, fn RouteHandler) {

}

// After every methods
func After(pattern string, fn RouteHandler) {

}

func AfterAll(fn RouteHandler) {

}

// Handling Exceptions
func Error(err error, fn RouteHandler) {

}