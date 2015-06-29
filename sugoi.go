package main
import (
	"log"
	"net/http"
	"errors"
	"regexp"
	"strconv"
	"fmt"
	"encoding/json"
)

type Response interface{}
type RouteHandler func(Request) Response

type Person struct {
	name 	string  `json:"name"`
	age 	int     `json:"age"`
}

func main() {
	s := NewSugoi("8085")

	s.GET("/", func(req Request) Response {
		return "homepage"
	})

	s.GET("/notfound", func(req Request) Response {
		return NotFound()
	})

	s.GET("/hello/:name", func(req Request) Response {
		name := req.GetAttribute("name")

		return "hello " + name
	})

	s.GET("/person", func(req Request) Response {

		person := &Person {
			name: "Zoob",
			age: 38,
		}

		return person
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

func NotFound(msg ... string) HttpCode {
	var content string

	if len(msg) > 0 {
		content = msg[0]
	} else {
		content = "404 - Not Found"
	}

	return HttpCode{
		code: http.StatusNotFound,
		content: content,
	}
}

func NewSugoi(port string) *SugoiServer {

	return &SugoiServer{
		handler: NewWrappedHandler(),
		port: port,
	}
}

type SugoiServer struct {
	handler 	*WrappedHandler
	port 		string
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

	err := http.ListenAndServe(":" + s.port, s.handler)
	if err != nil {
		log.Fatal(err)
	}
}

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
	// TODO: Match BeforeAll

	// TODO: Match Before "pattern"

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

	// TODO: Match AfterAll

	// TODO: Match After "pattern"
}

type HttpCode struct {
	code 	int
	content string
}

func (h *HttpCode) GetCode() int {
	return h.code
}

func (h *HttpCode) GetContent() string {
	return h.content
}

func SendHttpResponse(response interface{}, w http.ResponseWriter) {
	ResponseHandler(response, w)
}

func SendHttpCodeResponse(codeResponse HttpCode, w http.ResponseWriter) {
	w.WriteHeader(codeResponse.GetCode())

	if codeResponse.GetContent() != "" {
		w.Write([]byte(codeResponse.GetContent()))
	}
}

func ResponseHandler(response interface{}, w http.ResponseWriter) {
	if val, ok := response.(string); ok {
		w.Write([]byte(val))
	} else
	if val, ok := response.(int); ok {
		log.Println("int", val)
	} else
	if val, ok := response.(HttpCode); ok {
		SendHttpCodeResponse(val, w)
	} else {
		log.Println("Handling Object >> JSON", response)
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			errorHttpCode := HttpCode{
				code: 500,
				content: "An error occured processing request",
			}
			SendHttpCodeResponse(errorHttpCode, w)
		}
	}
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

// Return Content Types
	// text
	// object : implicit json
	// various error code objects
	// specific content type (with converters)
		// xml, json, plain text, binary etc
	// with html template