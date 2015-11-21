package sugoi

import (
	"log"
	"net/http"
	"strings"
)


type HttpServer interface {
	SetStatic(url, dir string)
	GetRoutes(method string) []*Route

	Handle(method string, path string, fn RouteHandler) *Route
	Get(path string, fn RouteHandler) *Route
	Delete(path string, fn RouteHandler) *Route
	Put(path string, fn RouteHandler) *Route
	Post(path string, fn RouteHandler) *Route
	Options(path string, fn RouteHandler) *Route
	Patch(path string, fn RouteHandler) *Route

	Start()
	Stop()

	Before(fn PreFilter)
	Error(fn ErrorHandler)
}

func NewHttpServer(port string) HttpServer {
	return NewSugoi(port)
}

func NewSugoi(host string) *SugoiServer {
	return &SugoiServer{
		handler: NewWrappedHandler(),
		host:    host,
	}
}

type SugoiServer struct {
	handler *WrappedHandler
	host    string
}

func (s *SugoiServer) SetStatic(url, dir string) {
	s.handler.staticUrl = url
	s.handler.staticDir = dir
}

func (s *SugoiServer) GetRoutes(method string) []*Route {
	routes := []*Route{}

	for _, o := range s.handler.routes {
		if o.method == method {
			routes = append(routes, o)
		}
	}

	return routes
}

func (s *SugoiServer) add(method string, path string, fn RouteHandler) *Route {
	route := CreateNewRoute(path, method, fn)
	s.handler.routes = append(s.handler.routes, route)

	return route
}

func (s *SugoiServer) Handle(method string, path string, fn RouteHandler) *Route {
	return s.add(method, path, fn)
}

func (s *SugoiServer) Get(path string, fn RouteHandler) *Route {
	return s.Handle("get", path, fn)
}

func (s *SugoiServer) Delete(path string, fn RouteHandler) *Route {
	return s.Handle("delete", path, fn)
}

func (s *SugoiServer) Put(path string, fn RouteHandler) *Route {
	return s.Handle("put", path, fn)
}

func (s *SugoiServer) Post(path string, fn RouteHandler) *Route {
	return s.Handle("post", path, fn)
}

func (s *SugoiServer) Options(path string, fn RouteHandler) *Route {
	return s.Handle("options", path, fn)
}

func (s *SugoiServer) Patch(path string, fn RouteHandler) *Route {
	return s.Handle("patch", path, fn)
}

func Set404Page(s *SugoiServer, fn RouteHandler) {
	putDefaultHandler(s, http.StatusNotFound, fn)
}

func Set500Page(s *SugoiServer, fn RouteHandler) {
	putDefaultHandler(s, http.StatusInternalServerError, fn)
}

func Set401Page(s *SugoiServer, fn RouteHandler) {
	putDefaultHandler(s, http.StatusUnauthorized, fn)
}

func Set403Page(s *SugoiServer, fn RouteHandler) {
	putDefaultHandler(s, http.StatusForbidden, fn)
}

func putDefaultHandler(s *SugoiServer, code int, fn RouteHandler) {
	s.handler.defaultHandlers[code] = fn
}

func (s *SugoiServer) Start() {
	host := s.host
	if !strings.Contains(s.host, ":") {
		host = ":" + s.host
	}
	
	log.Println("Started HTTP Server ", host)
	err := http.ListenAndServe(host, s.handler)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SugoiServer) Stop() {

}

// Before every method
func (s *SugoiServer) Before(fn PreFilter) {
	s.handler.preFilters = append(s.handler.preFilters, fn)
}

// Handling Exceptions
func (s *SugoiServer) Error(fn ErrorHandler) {
	s.handler.errorHandlers = append(s.handler.errorHandlers, fn)
}
